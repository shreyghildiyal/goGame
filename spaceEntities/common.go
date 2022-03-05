package spaceEntities

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	config "github.com/shreyghildiyal/goGame/configs"
	"github.com/shreyghildiyal/goGame/gametext"
	imageutils "github.com/shreyghildiyal/goGame/imageUtils"
	"golang.org/x/image/font"
)

type SpaceEntity interface {
	GetDisplay() *imageutils.DispDetails
	GetCoordinates() (float64, float64)
	GetTextPosition(font font.Face) (int, int)
	GetName() string
}

func DrawSpaceEntity(screen *ebiten.Image, camX, camY, camZoom float64, entity SpaceEntity) {

	disp := entity.GetDisplay()

	x, y := disp.BaseWidth, disp.BaseHeight
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(x/2), -float64(y/2))
	// op.GeoM.Rotate(star.rotation)
	// op.GeoM.Translate(float64(x/2), float64(y/2))
	op.GeoM.Scale(disp.ScaleX*camZoom, disp.ScaleY*camZoom)
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

func GetTextPosition(f font.Face, entity SpaceEntity) (float64, float64) {
	return entity.GetCoordinates()
}
