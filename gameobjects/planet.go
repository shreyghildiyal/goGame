package gameobjects

import (
	"math"

	config "github.com/shreyghildiyal/goGame/configs"
	"github.com/shreyghildiyal/goGame/drawing"
	imageutils "github.com/shreyghildiyal/goGame/imageUtils"
)

type Planet struct {
	Id         string
	Name       string
	StarId     string
	PlanetType string
	SysX       float64
	SysY       float64
	SysRad     float64
}

type PlanetSaveObj struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	StarId     string `json:"starId"`
	PlanetType string `json:"planetType"`
	Disp       struct {
		X      float64 `json:"x"`
		Y      float64 `json:"y"`
		Height float64 `json:"height"`
		Width  float64 `json:"width"`
	} `json:"disp"`
}

func (pso PlanetSaveObj) ToPlanet(conf config.Configuration) (Planet, error) {

	pl := Planet{}
	pl.Id = pso.Id
	pl.Name = pso.Name
	pl.PlanetType = pso.PlanetType
	pl.SysX = pso.Disp.X
	pl.SysY = pso.Disp.Y
	pl.SysRad = math.Sqrt((pl.SysX * pl.SysX) + (pl.SysY * pl.SysY))
	return pl, nil
}

func (pso PlanetSaveObj) GetSystemDrawables(conf config.Configuration) ([]drawing.Drawable, error) {
	systemImage, err := imageutils.GetImageFromMap(pso.PlanetType)
	if err != nil {
		return nil, nil
	}
	systemSprite := drawing.Drawable{
		X:            pso.Disp.X,
		Y:            pso.Disp.Y,
		Image:        systemImage,
		TargetHeight: pso.Disp.Height,
		TargetWidth:  pso.Disp.Width,
		RotAngle:     0,
	}
	err = validateSprite(systemSprite)
	if err != nil {
		return nil, err
	}

	ret := []drawing.Drawable{systemSprite}
	return ret, nil
}
