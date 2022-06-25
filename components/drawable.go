package components

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Drawable struct {
	TargetHeight float64 `json:"baseHeight"`
	TargetWidth  float64 `json:"baseWidth"`
	ScaleX       float64
	ScaleY       float64
	Image        *ebiten.Image
	ZLevel       int
	baseComponent
}
