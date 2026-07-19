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

	// fmt.Printf("cameraX: %.3f, cameraY: %.3f, taregtHeight: %.3f, targetWidth: %.3f, X: %.3f, Y: %.3f\n", camera.X, camera.Y, d.TargetHeight, d.TargetWidth, d.X, d.Y)

	// geom := ebiten.GeoM{}
	// // rotate the image. For this we need to move the center of the image to 0,0. we do this before scaling and translation as its independent of zoom level
	// geom.Translate(-float64(d.Image.Bounds().Dx())/2, -float64(d.Image.Bounds().Dy())/2)
	// geom.Rotate(d.RotAngle)
	// // scale the image to desired size.
	// geom.Scale(camera.Zoom*d.TargetWidth/float64(d.Image.Bounds().Dx()), camera.Zoom*d.TargetHeight/float64(d.Image.Bounds().Dy()))
	// // since we have already translated the image to 0,0 earlier, we dont need to cater for image size
	// // we need to translate using three coordinates. the object location, camera location and screen center.
	// // object at 0,0 and camera at 0,0 will be displayed at screen center
	// // object at 10,0 and camera at 0,0 will be displayed at screen center plus 10,0
	// // object at 0,0 and camera at 10,0 will be displayed at screen center plus -10,0
	// geom.Translate(float64(screen.Bounds().Dx())/2, float64(screen.Bounds().Dy())/2)
	// geom.Translate(camera.Zoom*(d.X-camera.X), camera.Zoom*(d.Y+camera.Y))
	geom := d.GetGeometry(float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy()), camera)

	opts := ebiten.DrawImageOptions{
		GeoM: geom,
	}

	screen.DrawImage(d.Image, &opts)
}

func (d *Drawable) GetGeometry(screenWidth, screenHeight float64, camera camera.Camera) ebiten.GeoM {

	geom := ebiten.GeoM{}
	// rotate the image. For this we need to move the center of the image to 0,0. we do this before scaling and translation as its independent of zoom level
	geom.Translate(-float64(d.Image.Bounds().Dx())/2, -float64(d.Image.Bounds().Dy())/2)
	geom.Rotate(d.RotAngle)
	// scale the image to desired size.
	geom.Scale(camera.Zoom*d.TargetWidth/float64(d.Image.Bounds().Dx()), camera.Zoom*d.TargetHeight/float64(d.Image.Bounds().Dy()))
	// since we have already translated the image to 0,0 earlier, we dont need to cater for image size
	// we need to translate using three coordinates. the object location, camera location and screen center.
	// object at 0,0 and camera at 0,0 will be displayed at screen center
	// object at 10,0 and camera at 0,0 will be displayed at screen center plus 10,0
	// object at 0,0 and camera at 10,0 will be displayed at screen center plus -10,0
	geom.Translate(screenWidth/2, screenHeight/2)
	geom.Translate(camera.Zoom*(d.X-camera.X), camera.Zoom*(d.Y-camera.Y))

	return geom
}
