package drawing

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/shreyghildiyal/goGame/camera"
)

type Drawable struct {
	X            float64
	Y            float64
	TargetHeight float64
	TargetWidth  float64
	Image        *ebiten.Image
}

func (d *Drawable) Draw(screen *ebiten.Image, camera camera.Camera) {

	geom := ebiten.GeoM{}
	geom.Scale(camera.Zoom*d.TargetWidth/float64(d.Image.Bounds().Dx()), camera.Zoom*d.TargetHeight/float64(d.Image.Bounds().Dy()))
	geom.Translate(camera.Zoom*(camera.X-(d.X-(d.TargetWidth/2))), camera.Zoom*(camera.Y-(d.Y-(d.TargetHeight/2))))

	opts := ebiten.DrawImageOptions{
		GeoM: ebiten.GeoM{},
	}

	screen.DrawImage(d.Image, &opts)
}
