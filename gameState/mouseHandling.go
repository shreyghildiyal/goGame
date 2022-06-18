package gamestate

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/shreyghildiyal/goGame/spaceEntities"
)

func HandleMouseInput(dt time.Duration, g *Game) {

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if g.CurrentView == GalaxyView {
			fmt.Println("Just licked left mouse button")
			for _, system := range g.Systems {
				cursorX, cursorY := ebiten.CursorPosition()
				if spaceEntities.IsClicked(system, float64(cursorX), float64(cursorY), g.Camera.X, g.Camera.Y, g.Camera.Zoom) {
					fmt.Println("System", system.Name, "was clicked")
					// set view to system view
					g.CurrentView = SystemView
					// set current system to the system in question
					g.CurrentSystemID = system.Id
				}
			}
		}
	}
}
