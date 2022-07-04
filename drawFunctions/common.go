package drawfunctions

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/shreyghildiyal/goGame/camera"
	"github.com/shreyghildiyal/goGame/components"
	"github.com/shreyghildiyal/goGame/gametext"
	"golang.org/x/image/font"
)

func DrawSprite(screen *ebiten.Image, camera *camera.Camera, drawable *components.Drawable, coordinates components.Coordinates) {

	// disp := entity.GetDisplay()

	x, y := drawable.TargetHeight, drawable.TargetHeight
	op := &ebiten.DrawImageOptions{}

	// op.GeoM.Rotate(star.rotation)
	// op.GeoM.Translate(float64(x/2), float64(y/2))
	op.GeoM.Scale(drawable.ScaleX*camera.Zoom, drawable.ScaleY*camera.Zoom)
	op.GeoM.Translate(-float64(x/2), -float64(y/2))
	op.GeoM.Translate(coordinates.AsFloatPair())
	op.GeoM.Translate(camera.X, camera.Y)

	// textX, textY := GetTextPosition(gametext.SpaceDisplayFont, coordinates.AsFloatPair())

	if drawable.Image == nil {
		// fmt.Println("image from", entity.GetType(), "was nil somehow")
	} else if gametext.SpaceDisplayFont == nil {
		fmt.Println("font is nil")
	} else {
		screen.DrawImage(drawable.Image, op)
		// text.Draw(screen, entity.GetName(), gametext.SpaceDisplayFont, int(textX+camX), int(textY+camY), config.GetConfig().Text.Colour)
	}
}

func GetTextPosition(f font.Face, x, y float64) (float64, float64) {
	return x, y
}
