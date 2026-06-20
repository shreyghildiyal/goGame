package drawing

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/shreyghildiyal/goGame/camera"
)

type Drawable struct {
	X            float64 // center of image/entity
	Y            float64 // center of image/entity
	TargetHeight float64
	TargetWidth  float64
	Image        *ebiten.Image
	RotAngle     float64 // angle in radians to rotate by
}

func (d *Drawable) Draw(screen *ebiten.Image, camera camera.Camera) {

	geom := ebiten.GeoM{}
	// rotate the image. For this we need to move the center of the image to 0,0. we do this before scaling and translation as its independent of zoom level
	geom.Translate(-float64(d.Image.Bounds().Dx())/2, -float64(d.Image.Bounds().Dy())/2)
	geom.Rotate(d.RotAngle)
	// scale the image to desired size.
	geom.Scale(camera.Zoom*d.TargetWidth/float64(d.Image.Bounds().Dx()), camera.Zoom*d.TargetHeight/float64(d.Image.Bounds().Dy()))
	// since we have already translated the image to 0,0 earlier, we dont need to cater for image size
	geom.Translate(camera.Zoom*(camera.X-d.X), camera.Zoom*(camera.Y-d.Y))

	opts := ebiten.DrawImageOptions{
		GeoM: ebiten.GeoM{},
	}

	screen.DrawImage(d.Image, &opts)
}
