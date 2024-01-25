package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Exporter
type RunType string

const (
	IN_SERVICE RunType = "in-service"
	JS         RunType = "js"
)

type Exporter struct {
	ID                   string  `bson:"_id,omitempty"`
	Name                 string  `bson:"name"`
	Description          string  `bson:"description"`
	Dimensions           []int   `bson:"dimensions"`
	Type                 RunType `bson:"type"`
	ExportScriptLocation *string `bson:"location"`
}

// Model
type Model struct {
	ID               string `json:"_id" bson:"_id"`
	Name             string `json:"name" bson:"name"`
	Description      string `json:"description" bson:"description"`
	Dims             []int  `json:"dims" bson:"dims"`
	TensorflowLayers bson.D `json:"tfLayers" bson:"tfLayers"`
}

// TrainingRun
type TrainingRun struct {
	ID              string  `json:"_id,omitempty"`
	ModelId         string  `json:"modelId"`
	Name            string  `json:"name"`
	DataSet         string  `json:"dataSet"`
	Time            string  `json:"time"`
	F1Train         float64 `json:"f1Train"`
	F1Test          float64 `json:"f1Test"`
	TrainingHistory bson.M  `json:"trainingHistory"`
	BatchSize       int     `json:"batchSize"`
	Epochs          int     `json:"epochs"`
}

// User
type Role string

const (
	ADMIN  Role = "admin"
	CLIENT Role = "client"
)

type UserData struct {
	ID    string `bson:"_id,omitempty"`
	Role  Role   `bson:"role"`
	Email string `bson:"email"`
	Key   string `bson:"key"`
}

// Request
type RequestDataLabel struct {
	IsLabeled bool   `json:"isLabeled"  validate:"required"`
	Blocklist string `json:"blocklist"  validate:"required"`
}

type RequestDataResponse struct {
	DocumentId        string              `json:"documentId"`
	DocumentLifecycle string              `json:"documentLifecycle"`
	FrameId           int                 `json:"frameId"`
	FrameType         string              `json:"frameType"`
	FromCache         bool                `json:"fromCache"`
	Initiator         string              `json:"initiator"`
	Ip                string              `json:"ip"`
	Method            string              `json:"method"`
	ParentFrameId     int                 `json:"parentFrameId"`
	RequestId         string              `json:"requestId"`
	RequestHeaders    []map[string]string `json:"responseHeaders"`
	StatusCode        int                 `json:"statusCode"`
	StatusLine        string              `json:"statusLine"`
	TabId             int                 `json:"tabId"`
	TimeStamp         float32             `json:"timeStamp"`
	Type              string              `json:"type"`
	URL               string              `json:"url"`
}

type RequestData struct {
	ID                string              `json:"_id" bson:"_id,omitempty"`
	DocumentId        string              `json:"documentId" bson:"documentId"`
	DocumentLifecycle string              `json:"documentLifecycle" bson:"documentLifecycle"`
	FrameId           int                 `json:"frameId" bson:"frameId"`
	FrameType         string              `json:"frameType" bson:"frameType"`
	Initiator         string              `json:"initiator" bson:"initiator"`
	Method            string              `json:"method" bson:"method"`
	ParentFrameId     int                 `json:"parentFrameId" bson:"parentFrameId"`
	RequestId         string              `json:"requestId" bson:"requestId"`
	TabId             int                 `json:"tabId" bson:"tabId"`
	TimeStamp         float32             `json:"timeStamp" bson:"timeStamp"`
	Type              string              `json:"type" bson:"type"`
	URL               string              `json:"url" bson:"url" validate:"required"`
	RequestHeaders    []map[string]string `json:"requestHeaders" bson:"requestHeaders"`
	Response          RequestDataResponse `json:"response" bson:"response"`
	Success           bool                `json:"success" bson:"success"`
	Dataset           string              `json:"dataset" bson:"dataset"`
	Labels            []RequestDataLabel  `json:"labels" bson:"labels" validate:"required"`
}

type ReducedRequestData struct {
	DocumentId        string              `json:"documentId" bson:"documentId"`
	DocumentLifecycle string              `json:"documentLifecycle" bson:"documentLifecycle"`
	FrameId           int                 `json:"frameId" bson:"frameId"`
	FrameType         string              `json:"frameType" bson:"frameType"`
	Initiator         string              `json:"initiator" bson:"initiator"`
	Method            string              `json:"method" bson:"method"`
	ParentFrameId     int                 `json:"parentFrameId" bson:"parentFrameId"`
	RequestId         string              `json:"requestId" bson:"requestId"`
	TabId             int                 `json:"tabId" bson:"tabId"`
	TimeStamp         float32             `json:"timeStamp" bson:"timeStamp"`
	Type              string              `json:"type" bson:"type"`
	URL               string              `json:"url" bson:"url" validate:"required"`
	RequestHeaders    []map[string]string `json:"requestHeaders" bson:"requestHeaders"`
	Response          RequestDataResponse `json:"response" bson:"response"`
	Success           bool                `json:"success" bson:"success"`
	Tracker           bool                `json:"tracker" bson:"tracker"`
}

// Export Metrics
type ExportMetrics struct {
	Tracker    int    `json:"tracker" bson:"tracker"`
	NonTracker int    `json:"nonTracker" bson:"nonTracker"`
	Total      int    `json:"total" bson:"total"`
	Error      string `json:"error" bson:"error"`
}

type ExportRun struct {
	ID           string         `json:"_id" bson:"_id,omitempty"`
	ExporterId   string         `json:"exporterId" bson:"exporterId"`
	ExporterName string         `json:"exporterName" bson:"exporterName"`
	Metrics      *ExportMetrics `json:"metrics" bson:"metrics"`
	Start        time.Time      `json:"start" bson:"start"`
	End          time.Time      `json:"end" bson:"end"`
}
