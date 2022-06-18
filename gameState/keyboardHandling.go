package gamestate

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	config "github.com/shreyghildiyal/goGame/configs"
)

func HandleKeyboardInput(dt time.Duration, g *Game) {
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
}
