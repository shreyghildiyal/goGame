package gameobjects

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/shreyghildiyal/goGame/camera"
	config "github.com/shreyghildiyal/goGame/configs"
	"github.com/shreyghildiyal/goGame/drawing"
	imageutils "github.com/shreyghildiyal/goGame/imageUtils"
)

type Star struct {
	// drawing.Drawable

	Id           string
	Name         string
	SystemRadius float64
	GalX         float64
	GalY         float64
	Planets      []string
}

type StarSaveObj struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	StarType   string `json:"starType"`
	GalaxyDisp struct {
		X      float64 `json:"x"`
		Y      float64 `json:"y"`
		Height float64 `json:"height"`
		Width  float64 `json:"width"`
	} `json:"galaxyDisp"`
	SystemDisp struct {
		Height       float64 `json:"height"`
		Width        float64 `json:"width"`
		SystemRadius float64 `json:"systemRadius"`
	} `json:"systemDisp"`
}

func (sso StarSaveObj) GetSystemDrawables(conf config.Configuration) ([]drawing.Drawable, error) {
	systemImage, err := imageutils.GetImageFromMap(sso.StarType)
	if err != nil {
		return nil, nil
	}
	systemSprite := drawing.Drawable{
		X:            0,
		Y:            0,
		Image:        systemImage,
		TargetHeight: sso.SystemDisp.Height,
		TargetWidth:  sso.SystemDisp.Width,
		RotAngle:     0,
	}
	err = validateSprite(systemSprite)
	if err != nil {
		return nil, err
	}

	ret := []drawing.Drawable{systemSprite}
	return ret, nil
}

func validateSprite(d drawing.Drawable) error {
	if d.TargetHeight < 1 || d.TargetWidth < 1 {
		return fmt.Errorf("the sprite doesnt seem to have a valid size")
	}
	return nil
}

func (sso StarSaveObj) GetGalaxyDrawables(conf config.Configuration) ([]drawing.Drawable, error) {

	galaxyImage, err := imageutils.GetImageFromMap(sso.StarType)
	if err != nil {
		return nil, err
	}
	galSprite := drawing.Drawable{
		X:            sso.GalaxyDisp.X,
		Y:            sso.GalaxyDisp.Y,
		Image:        galaxyImage,
		TargetHeight: sso.GalaxyDisp.Height,
		TargetWidth:  sso.GalaxyDisp.Width,
		RotAngle:     0,
	}

	err = validateSprite(galSprite)
	if err != nil {
		return nil, err
	}

	ret := []drawing.Drawable{galSprite}
	return ret, nil
}

func (sso StarSaveObj) ToStar(conf config.Configuration) (Star, error) {

	star := Star{
		Name:         sso.Name,
		Id:           sso.Id,
		SystemRadius: sso.SystemDisp.SystemRadius,
		GalX:         sso.GalaxyDisp.X,
		GalY:         sso.GalaxyDisp.Y,
		Planets:      []string{},
	}
	return star, nil
}

func NewStar(id string, name string, galSprite drawing.Drawable, systemSprite drawing.Drawable) Star {

	s := Star{}
	s.Id = id
	s.Name = name
	return s
}

func (s Star) DrawSystemBoundary(screen *ebiten.Image, camera camera.Camera) {
	panic("unimplemented")
}
