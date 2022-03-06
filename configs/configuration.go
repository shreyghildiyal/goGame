package config

import (
	"encoding/json"
	"image/color"
	"io/ioutil"
	"log"
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

type Configuration struct {
	BackgroundImagePath   string            `json:"backgroundImage"`
	PlanetsFile           string            `json:"planetsFile"`
	SystemsFile           string            `json:"systemsFile"`
	SystemsConnectionFile string            `json:"systemsConnectionFile"`
	Text                  TextConf          `json:"text"`
	PlanetImages          map[string]string `json:"planetImages"`
	StarImages            map[string]string `json:"starImages"`
	Camera                Camera            `json:"camera"`
	WarpLines             WarpLines         `json:"warpLines"`
}

type WarpLines struct {
	Colour color.RGBA `json:"colour"`
}

func GetConfig() Configuration {
	if config == nil {
		confMutex.Lock()
		if config == nil {
			LoadConfig()
		}
		confMutex.Unlock()
	}
	return *config
}

func LoadConfig() {
	confFile := "conf.json"

	data, err := ioutil.ReadFile(confFile)
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
