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
	x, y := d.Image.Size()
	fmt.Println("image size", x, y)
	d.ScaleX = float64(d.BaseWidth) / float64(x)
	d.ScaleY = float64(d.BaseHeight) / float64(y)
	fmt.Println("Image scale", d.ScaleX, d.ScaleY)
}
