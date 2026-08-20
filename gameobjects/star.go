package gameobjects

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/shreyghildiyal/goGame/camera"
	config "github.com/shreyghildiyal/goGame/configs"
	"github.com/shreyghildiyal/goGame/drawing"
	"github.com/shreyghildiyal/goGame/gametext"
	imageutils "github.com/shreyghildiyal/goGame/imageUtils"
)

type Star struct {
	// drawing.Drawable
	GalaxySprite drawing.Drawable
	SystemSprite drawing.Drawable
	Id           string `json:"id"`
	Name         string `json:"name"`
	SystemRadius float64
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

func validateSprite(d drawing.Drawable) error {
	if d.TargetHeight < 1 || d.TargetWidth < 1 {
		return fmt.Errorf("the sprite doesnt seem to have a valid size")
	}
	return nil
}

func (sso StarSaveObj) ToStar(conf config.Configuration) (Star, error) {
	// TODO: these can and should be different images in the future. also we need to check for errors
	galaxyImage, err := imageutils.GetImageFromMap(sso.StarType)
	if err != nil {
		return Star{}, nil
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
		return Star{}, err
	}
	systemImage, err := imageutils.GetImageFromMap(sso.StarType)
	if err != nil {
		return Star{}, nil
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
		return Star{}, err
	}
	star := Star{
		Name:         sso.Name,
		Id:           sso.Id,
		GalaxySprite: galSprite,
		SystemSprite: systemSprite,
		SystemRadius: sso.SystemDisp.SystemRadius,
	}
	return star, nil
}

func (s Star) DrawGalaxySprite(screen *ebiten.Image, camera camera.Camera) {
	s.GalaxySprite.Draw(screen, camera)
	s.GalaxySprite.DrawName(s.Name, gametext.SpaceDisplayFont, gametext.SpaceColour, screen, camera)
	// panic("unimplemented")
}

func NewStar(id string, name string, galSprite drawing.Drawable, systemSprite drawing.Drawable) Star {

	s := Star{}
	s.Id = id
	s.GalaxySprite = galSprite
	s.SystemSprite = systemSprite
	s.Name = name
	return s
}

func (s Star) DrawSystemBoundary(screen *ebiten.Image, camera camera.Camera) {
	panic("unimplemented")
}
