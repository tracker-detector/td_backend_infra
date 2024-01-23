package consumer

import (
	"os"
	"os/signal"
	"syscall"
	"tds/shared/messages"
	"tds/shared/models"
	"tds/shared/service"

	log "github.com/sirupsen/logrus"

	"github.com/streadway/amqp"
)

type IConsumer interface {
	Consume()
}

type ExportMessageConsumer struct {
	con             *amqp.Connection
	ch              *amqp.Channel
	requestRepo     models.RequestRepository
	storageService  service.IStorageService
	exporterService service.IExporterService
}

func NewExportMessageConsumer(requestRepo models.RequestRepository, storageService service.IStorageService, exporterService service.IExporterService) *ExportMessageConsumer {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declare the exports queue
	_, err = ch.QueueDeclare(
		"exports", // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare an exports queue: %v", err)
	}
	return &ExportMessageConsumer{
		requestRepo:     requestRepo,
		storageService:  storageService,
		exporterService: exporterService,
		con:             conn,
		ch:              ch,
	}
}

func (c *ExportMessageConsumer) Consume() {
	msgs, err := c.ch.Consume(
		"exports", // queue name
		"",
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	log.Println("Export ConsumerService started. Waiting for messages...")

	go func() {
		for msg := range msgs {
			c.handleMessage(msg.Body)
		}
	}()

	<-stopChan
	log.Println("Shutting down Export ConsumerService.")
}

func (c *ExportMessageConsumer) handleMessage(msg []byte) {
	job, err := messages.DeserializeJob(string(msg))
	if err != nil {
		log.Errorf("Failed to deserialize job: %v", err)
		return
	}
	exporterId := job.Args[0]
	reducer := job.Args[1]

}