package imageutils

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type DispDetails struct {
	TargetHeight float64 `json:"baseHeight"`
	TargetWidth  float64 `json:"baseWidth"`
	ScaleX       float64
	ScaleY       float64
	Image        *ebiten.Image
}

func (d *DispDetails) InitDetails() {
	width, height := d.Image.Size()
	fmt.Println("image size", width, height)
	d.ScaleX = d.TargetWidth / float64(width)
	d.ScaleY = d.TargetHeight / float64(height)
	fmt.Println("Image scale", d.ScaleX, d.ScaleY)
}
