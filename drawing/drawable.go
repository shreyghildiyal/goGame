package drawing

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/shreyghildiyal/goGame/camera"
)

type Drawable struct {
	X            float64 `json:"x"` // center of image/entity
	Y            float64 `json:"y"` // center of image/entity
	TargetHeight float64 ``
	TargetWidth  float64
	Image        *ebiten.Image
	RotAngle     float64 // angle in radians to rotate by
}

func (d Drawable) DrawName(name string, face *text.GoTextFace, textColor color.Color, screen *ebiten.Image, camera camera.Camera) {
	// this geometry moves and scales the drawable image to center at X,Y in gamespace, with adjustments to camera location and zoom.
	geom := d.GetGeometry(float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy()), camera)

	// if we transform the mid point of lower edge using this geometry, we will have the location of the middle of the desired text
	midX, midY := geom.Apply(float64(d.Image.Bounds().Dx())/2, float64(d.Image.Bounds().Dy()))

	// we then find the width of the text using measure
	textWidth, _ := text.Measure(name, face, face.Size)

	// we then set the display location of the text to coord we have minus half of text width from measure
	textX := midX - (textWidth / 2)
	textY := midY
	// fmt.Println("name", name, "textWidth", textWidth, "textX", textX, "textY", textY, "d.X", d.X, "d.Y", d.Y, "imgWidht", d.Image.Bounds().Dx(), "imgHeight", d.Image.Bounds().Dy())

	op := &text.DrawOptions{}
	op.GeoM.Translate(textX, textY)
	op.ColorScale.ScaleWithColor(textColor)
	text.Draw(screen, name, face, op)
}

func (d *Drawable) Draw(screen *ebiten.Image, camera camera.Camera) {

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

func (d *Drawable) IsClicked(screenWidth, screenHeight float64, camera camera.Camera, mouseX, mouseY float64) bool {

	// the idea is to translate the corners of the drawable based on camera etc. then check if mouseX and mouseY belong to the dawable range or not

	geom := d.GetGeometry(screenWidth, screenHeight, camera)

	minX, minY := geom.Apply(0, 0)
	log.Println("min loc", minX, minY)
	if mouseX < minX || mouseY < minY {
		return false
	}
	bx := float64(d.Image.Bounds().Dx())
	by := float64(d.Image.Bounds().Dy())
	maxX, maxY := geom.Apply(bx, by)
	log.Println("max loc", maxX, maxY)
	if mouseX > maxX || mouseY > maxY {
		return false
	}

	return true
}
