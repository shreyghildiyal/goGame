package gamestate

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	config "github.com/shreyghildiyal/goGame/configs"
)

func HandleKeyboardInput(dt time.Duration, g *GameState) {

	var timeMulti float64 = float64(dt.Microseconds()) / 1000000
	// fmt.Println("time Multi", timeMulti)
	// fmt.Println("Speed", config.GetConfig().Camera.SpeedX, config.GetConfig().Camera.SpeedY)
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.Camera.Y += config.GetConfig().Camera.SpeedY * timeMulti
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.Camera.Y -= config.GetConfig().Camera.SpeedY * timeMulti
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.Camera.X += config.GetConfig().Camera.SpeedX * timeMulti
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.Camera.X -= config.GetConfig().Camera.SpeedX * timeMulti
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if g.CurrentView == SystemView {
			// set view to galaxy view
			g.CurrentView = GalaxyView
			// set camera location to system location in galaxy
			// g.Camera.X = g.Systems[g.CurrentSystemID].X
			// g.Camera.Y = g.Systems[g.CurrentSystemID].Y
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyS) {
		fmt.Println("Calling SaveGame")
		g.saveGame()
	}
}
