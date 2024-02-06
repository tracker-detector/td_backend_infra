package resources

import (
	"encoding/json"
	"os"

	"github.com/Tracking-Detector/td_backend_infra/public/models"
)

const (
	heroResourcePath     = "./assets/content/hero.json"
	featuresResourcePath = "./assets/content/features.json"
	productsResourcePath = "./assets/content/products.json"
)

func LoadHomeResource() *models.Home {
	home := models.NewHome()
	hero := &models.Hero{}
	loadResource(heroResourcePath, hero)
	home.Hero = hero
	features := &models.Features{}
	loadResource(featuresResourcePath, features)
	home.Features = features
	products := &models.Products{}
	loadResource(productsResourcePath, products)
	home.Products = products
	return home
}

func loadResource(filePath string, resource interface{}) {
	// Read file path and map to models.Hero
	file, err := os.Open(heroResourcePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(resource)
	if err != nil {
		panic(err)
	}
}
