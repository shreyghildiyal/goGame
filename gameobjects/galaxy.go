package gameobjects

import (
	"encoding/json"
	"fmt"
	"os"

	config "github.com/shreyghildiyal/goGame/configs"
)

type Galaxy struct {
	Stars   map[int]Star   `json:"stars"`
	Planets map[int]Planet `json:"planets"`
}

type GalaxySaveObj struct {
	Stars   map[int]StarSaveObj   `json:"stars"`
	Planets map[int]PlanetSaveObj `json:"planets"`
}

func LoadGalaxy(saveFile string, conf config.Configuration) (Galaxy, error) {

	fileData, err := os.ReadFile(saveFile)
	if err != nil {
		return Galaxy{}, err
	}
	saveGal := GalaxySaveObj{}
	err = json.Unmarshal(fileData, &saveGal)
	if err != nil {
		return Galaxy{}, err
	}

	gal := Galaxy{
		Stars:   map[int]Star{},
		Planets: map[int]Planet{},
	}

	if len(saveGal.Stars) < 1 {
		return Galaxy{}, fmt.Errorf("The save file had no stars in it")
	}

	for id, st := range saveGal.Stars {
		gal.Stars[id], err = st.ToStar(conf)
		if err != nil {
			return Galaxy{}, err
		}
	}

	if len(saveGal.Planets) < 1 {
		return Galaxy{}, fmt.Errorf("The save file had no planets in it")
	}

	for id, st := range saveGal.Planets {
		gal.Planets[id], err = st.ToPlanet(conf)
		if err != nil {
			return Galaxy{}, err
		}
	}

	return gal, nil
}
