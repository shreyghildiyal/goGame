package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

var config *Configuration
var confMutex sync.Mutex

type Configuration struct {
	BackgroundImagePath string `json:"backgroundImage"`
	PlanetsFile         string `json:"planetsFile"`
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
