package gametext

import (
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	config "github.com/shreyghildiyal/goGame/configs"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	SpaceDisplayFont font.Face
)

func InitFonts() {

	tt, _ := opentype.Parse(fonts.MPlus1pRegular_ttf)
	SpaceDisplayFont, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(config.GetConfig().Text.Size),
		DPI:     config.GetConfig().Text.Dpi,
		Hinting: font.HintingFull,
	})
}
