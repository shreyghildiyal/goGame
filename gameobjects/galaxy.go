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
	Stars map[int]StarSaveObj `json:"stars"`
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
		Stars: map[int]Star{},
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

	return gal, nil
}
