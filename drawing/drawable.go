package drawing

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Drawable struct {
	X            float64
	Y            float64
	TargetHeight float64
	TargetWidth  float64
	Image        *ebiten.Image
}

func (d *Drawable) Draw(screen *ebiten.Image) {

	geom := ebiten.GeoM{}
	geom.Scale(d.TargetWidth/float64(d.Image.Bounds().Dx()), d.TargetHeight/float64(d.Image.Bounds().Dy()))
	geom.Translate(d.X-(d.TargetWidth/2), d.Y-(d.TargetHeight/2))

	opts := ebiten.DrawImageOptions{
		GeoM: ebiten.GeoM{},
	}

	screen.DrawImage(d.Image, &opts)
}
