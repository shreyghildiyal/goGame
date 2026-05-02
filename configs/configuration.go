package config

import (
	"encoding/json"
	"image/color"
	"log"
	"os"
	"sync"
)

var config *Configuration
var confMutex sync.Mutex

type TextConf struct {
	Colour color.RGBA `json:"colour"`
	Size   int        `json:"size"`
	Dpi    float64    `json:"dpi"`
}

type Camera struct {
	SpeedX    float64 `json:"SpeedX"`
	SpeedY    float64 `json:"SpeedY"`
	SpeedZoom float64 `json:"SpeedZoom"`
}

type ScreenSize struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

type Configuration struct {
	Text         TextConf      `json:"text"`
	ImagesFile   string        `json:"imagesFile"`
	Camera       Camera        `json:"camera"`
	ScreenSize   ScreenSize    `json:"screenSize"`
	EntitiesFile string        `json:"entitiesFile"`
	SaveGameDir  string        `json:"saveGameDir"`
	MarketConfs  MarketConfigs `json:"marketConfigs"`
}

type MarketConfigs struct {
	PriceAdjustmentFactor    float64 `json:"priceAdjustmentFactor"`
	AdjustmentMinBuyQuantity float64 `json:"adjustmentMinBuyQuantity"`
}

type WarpLines struct {
	Colour color.RGBA `json:"colour"`
}

func GetConfig() Configuration {
	if config == nil {
		confMutex.Lock()
		if config == nil {
			loadConfig()
		}
		confMutex.Unlock()
	}
	return *config
}

func loadConfig() {
	confFile := "conf.json"

	data, err := os.ReadFile(confFile)
	if err != nil {
		log.Fatal(err)
	}

	newConf := Configuration{}

	err = json.Unmarshal(data, &newConf)
	if err != nil {
		log.Fatal(err)
	}

	config = &newConf
}
