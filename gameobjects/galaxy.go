package gameobjects

import (
	"encoding/json"
	"fmt"
	"os"

	config "github.com/shreyghildiyal/goGame/configs"
	"github.com/shreyghildiyal/goGame/drawing"
)

type Galaxy struct {
	Stars   map[string]Star   `json:"stars"`
	Planets map[string]Planet `json:"planets"`
}

type GalaxySaveObj struct {
	Stars   map[string]StarSaveObj   `json:"stars"`
	Planets map[string]PlanetSaveObj `json:"planets"`
}

func LoadGalaxy(saveFile string, conf config.Configuration) (Galaxy, drawing.DrawableRegistry, error) {

	fileData, err := os.ReadFile(saveFile)
	if err != nil {
		return Galaxy{}, drawing.DrawableRegistry{}, err
	}
	saveGal := GalaxySaveObj{}
	err = json.Unmarshal(fileData, &saveGal)
	if err != nil {
		return Galaxy{}, drawing.DrawableRegistry{}, err
	}

	gal := Galaxy{
		Stars:   map[string]Star{},
		Planets: map[string]Planet{},
	}
	drawReg := drawing.NewDrawableRegistry()

	if len(saveGal.Stars) < 1 {
		return Galaxy{}, drawing.DrawableRegistry{}, fmt.Errorf("The save file had no stars in it")
	}

	for id, st := range saveGal.Stars {
		gal.Stars[id], err = st.ToStar(conf)
		if err != nil {
			return Galaxy{}, drawing.DrawableRegistry{}, err
		}
		galDrawables, err := st.GetGalaxyDrawables(conf)

		if err != nil {
			return Galaxy{}, drawing.DrawableRegistry{}, err
		}
		drawReg.AddGalaxyDrawables(st.Id, galDrawables)
		systemDrawables, err := st.GetSystemDrawables(conf)
		if err != nil {
			return Galaxy{}, drawing.DrawableRegistry{}, err
		}
		drawReg.AddSystemDrawables(st.Id, systemDrawables)
	}

	if len(saveGal.Planets) < 1 {
		return Galaxy{}, drawing.DrawableRegistry{}, fmt.Errorf("The save file had no planets in it")
	}

	for id, st := range saveGal.Planets {
		gal.Planets[id], err = st.ToPlanet(conf)
		if err != nil {
			return Galaxy{}, drawing.DrawableRegistry{}, err
		}
		// planetDrawables, err := st.
	}

	return gal, drawReg, nil
}
