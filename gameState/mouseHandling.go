package gamestate

import (
	"time"
)

func HandleMouseInput(dt time.Duration, g *GameState) {

	// if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
	// 	if g.CurrentView.ViewType == GalaxyView {
	// 		fmt.Println("Just licked left mouse button")
	// 		for _, system := range g.Systems {
	// 			cursorX, cursorY := ebiten.CursorPosition()
	// 			if spaceEntities.IsClicked(system, float64(cursorX), float64(cursorY), g.Camera.X, g.Camera.Y, g.Camera.Zoom) {
	// 				fmt.Println("System", system.Name, "was clicked")
	// 				// set view to system view
	// 				g.CurrentView = View{ViewType: SystemView}
	// 				// set current system to the system in question
	// 				g.CurrentSystem = system
	// 			}
	// 		}
	// 	}
	// }
}
