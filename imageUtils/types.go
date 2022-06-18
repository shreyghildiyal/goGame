package imageutils

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type DispDetails struct {
	BaseHeight int `json:"baseHeight"`
	BaseWidth  int `json:"baseWidth"`
	ScaleX     float64
	ScaleY     float64
	Image      *ebiten.Image
}

func (d *DispDetails) InitDetails() {
	width, height := d.Image.Size()
	fmt.Println("image size", width, height)
	d.ScaleX = float64(d.BaseWidth) / float64(width)
	d.ScaleY = float64(d.BaseHeight) / float64(height)
	fmt.Println("Image scale", d.ScaleX, d.ScaleY)
}
