package inputs

import "github.com/hajimehoshi/ebiten/v2"

type Mouse struct {
	X           int
	Y           int
	ButtonsDown map[ebiten.MouseButton]bool
	Dragging    map[ebiten.MouseButton]bool
}
