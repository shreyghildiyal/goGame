package gameobjects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/shreyghildiyal/goGame/drawing"
)

type Star struct {
	drawing.Drawable
}

func NewStar(x, y float64, id int, sprite *ebiten.Image) {

	s := Star{}
	s.X = x
	s.Y = y
	s.Drawable.Image = sprite
}
