package ui

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/shreyghildiyal/goGame/camera"
	"github.com/shreyghildiyal/goGame/gameobjects"
	"github.com/shreyghildiyal/goGame/gametext"
)

type GlobalUI struct {
}

func (g GlobalUI) Draw(screen *ebiten.Image, camera camera.Camera, galaxy gameobjects.Galaxy, tick int) {
	vector.FillRect(screen, 0, 0, float32(screen.Bounds().Dx()), 30, color.RGBA{30, 30, 40, 255}, true)

	op := &text.DrawOptions{}
	op.GeoM.Translate(5, 5)
	op.ColorScale.ScaleWithColor(color.RGBA{R: 255, G: 255, B: 255, A: 255})
	text.Draw(screen, fmt.Sprint(tick), gametext.SpaceDisplayFont, op)
}
