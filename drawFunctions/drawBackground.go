package drawfunctions

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func DrawBackground(screen *ebiten.Image, bg *ebiten.Image) {
	screen.DrawImage(bg, nil)
}
