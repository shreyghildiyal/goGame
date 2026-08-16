package gameobjects

import (
	"github.com/shreyghildiyal/goGame/drawing"
)

type Star struct {
	// drawing.Drawable
	GalaxySprite drawing.Drawable
	SystemSprite drawing.Drawable
	Id           int
	Name         string
}

func NewStar(id int, name string, galSprite drawing.Drawable, systemSprite drawing.Drawable) Star {

	s := Star{}
	s.Id = id
	s.GalaxySprite = galSprite
	s.SystemSprite = systemSprite

	return s
}
