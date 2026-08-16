package gameobjects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/shreyghildiyal/goGame/camera"
	"github.com/shreyghildiyal/goGame/drawing"
	"github.com/shreyghildiyal/goGame/gametext"
)

type Star struct {
	// drawing.Drawable
	GalaxySprite drawing.Drawable
	SystemSprite drawing.Drawable
	Id           int
	Name         string
}

func (s Star) DrawGalaxySprite(screen *ebiten.Image, camera camera.Camera) {
	s.GalaxySprite.Draw(screen, camera)
	s.GalaxySprite.DrawName(s.Name, gametext.SpaceDisplayFont, gametext.SpaceColour, screen, camera)
	// panic("unimplemented")
}

func NewStar(id int, name string, galSprite drawing.Drawable, systemSprite drawing.Drawable) Star {

	s := Star{}
	s.Id = id
	s.GalaxySprite = galSprite
	s.SystemSprite = systemSprite
	s.Name = name
	return s
}
