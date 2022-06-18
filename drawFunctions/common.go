package drawfunctions

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	config "github.com/shreyghildiyal/goGame/configs"
	"github.com/shreyghildiyal/goGame/gametext"
	"github.com/shreyghildiyal/goGame/spaceEntities"
	"golang.org/x/image/font"
)

func DrawSpaceEntity(screen *ebiten.Image, camX, camY, camZoom float64, entity spaceEntities.SpaceEntity) {

	disp := entity.GetDisplay()

	x, y := disp.BaseWidth, disp.BaseHeight
	op := &ebiten.DrawImageOptions{}

	// op.GeoM.Rotate(star.rotation)
	// op.GeoM.Translate(float64(x/2), float64(y/2))
	op.GeoM.Scale(disp.ScaleX*camZoom, disp.ScaleY*camZoom)
	op.GeoM.Translate(-float64(x/2), -float64(y/2))
	op.GeoM.Translate(entity.GetCoordinates())
	op.GeoM.Translate(camX, camY)

	textX, textY := GetTextPosition(gametext.SpaceDisplayFont, entity)

	if disp.Image == nil {
		// fmt.Println("image from", entity.GetType(), "was nil somehow")
	} else if gametext.SpaceDisplayFont == nil {
		fmt.Println("font is nil")
	} else {
		screen.DrawImage(disp.Image, op)
		text.Draw(screen, entity.GetName(), gametext.SpaceDisplayFont, int(textX+camX), int(textY+camY), config.GetConfig().Text.Colour)
	}
}

func GetTextPosition(f font.Face, entity spaceEntities.SpaceEntity) (float64, float64) {
	return entity.GetCoordinates()
}
