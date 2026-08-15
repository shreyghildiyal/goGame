package gamestate

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func HandleMouseInput(dt time.Duration, g *GameState) {

	// for _, star := range g.Galaxy.Stars {

	// }
	mouseX, mouseY := ebiten.CursorPosition()

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if g.CurrentView == GalaxyView {
			fmt.Println("Just licked left mouse button", mouseX, mouseY)
			for _, star := range g.Galaxy.Stars {
				// cursorX, cursorY := ebiten.CursorPosition()
				if star.GalaxySprite.IsClicked(float64(g.conf.ScreenSize.Width), float64(g.conf.ScreenSize.Height), g.Camera, float64(mouseX), float64(mouseY)) {
					fmt.Println("System", star.Id, "was clicked")
					// set view to system view
					g.CurrentView = SystemView
					// set current system to the system in question
					g.CurrentSystemId = star.Id
				}
			}
		}
	}
}
