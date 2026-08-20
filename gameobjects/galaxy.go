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

	gal, dr, err := loadStars(saveGal, gal, conf, drawReg)
	if err != nil {
		return gal, dr, err
	}

	if len(saveGal.Planets) < 1 {
		return Galaxy{}, drawing.DrawableRegistry{}, fmt.Errorf("The save file had no planets in it")
	}

	for id, pt := range saveGal.Planets {
		gal.Planets[id], err = pt.ToPlanet(conf)
		if err != nil {
			return Galaxy{}, drawing.DrawableRegistry{}, err
		}

		if star, found := gal.Stars[pt.StarId]; found {
			star.Planets = append(star.Planets, pt.Id)
			gal.Stars[pt.StarId] = star
		} else {
			return gal, drawReg, fmt.Errorf("No star with Id %s present. Planet with Id %s apparently belongs to it", pt.StarId, pt.Id)
		}

		sysDrawables, err := pt.GetSystemDrawables(conf)
		if err != nil {
			return Galaxy{}, drawing.DrawableRegistry{}, err
		}
		drawReg.AddSystemDrawables(pt.Id, sysDrawables)

	}

	return gal, drawReg, nil
}

func loadStars(saveGal GalaxySaveObj, gal Galaxy, conf config.Configuration, drawReg drawing.DrawableRegistry) (Galaxy, drawing.DrawableRegistry, error) {
	if len(saveGal.Stars) < 1 {
		return Galaxy{}, drawing.DrawableRegistry{}, fmt.Errorf("The save file had no stars in it")
	}

	for id, st := range saveGal.Stars {
		star, err := st.ToStar(conf)
		gal.Stars[id] = star
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
	return gal, drawReg, nil
}
