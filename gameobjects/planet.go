package gameobjects

import (
	config "github.com/shreyghildiyal/goGame/configs"
	"github.com/shreyghildiyal/goGame/drawing"
	imageutils "github.com/shreyghildiyal/goGame/imageUtils"
)

type Planet struct {
	Sprite     drawing.Drawable
	Id         int
	Name       string
	StarId     int
	PlanetType string
}

type PlanetSaveObj struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	StarId     int    `json:"starId"`
	PlanetType string `json:"planetType"`
	Disp       struct {
		X      float64 `json:"x"`
		Y      float64 `json:"y"`
		Height float64 `json:"height"`
		Width  float64 `json:"width"`
	} `json:"disp"`
}

func (p PlanetSaveObj) ToPlanet(conf config.Configuration) (Planet, error) {

	planetImage, err := imageutils.GetImageFromMap(p.PlanetType)
	if err != nil {
		return Planet{}, nil
	}
	planetSprite := drawing.Drawable{
		X:            p.Disp.X,
		Y:            p.Disp.Y,
		TargetHeight: p.Disp.Height,
		TargetWidth:  p.Disp.Width,
		Image:        planetImage,
	}

	pl := Planet{}
	pl.Id = p.Id
	pl.Name = p.Name
	pl.PlanetType = p.PlanetType
	pl.Sprite = planetSprite

	return pl, nil
}
