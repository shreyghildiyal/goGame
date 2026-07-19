package gameobjects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/shreyghildiyal/goGame/drawing"
)

type Star struct {
	drawing.Drawable
}

func NewStar(x, y, targetHeight, targetWidth, rot float64, id int, sprite *ebiten.Image) Star {

	s := Star{}
	s.X = x
	s.Y = y
	s.TargetHeight = targetHeight
	s.TargetWidth = targetWidth
	s.Drawable.Image = sprite
	s.RotAngle = rot

	return s
}
