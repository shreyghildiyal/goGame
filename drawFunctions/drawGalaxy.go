package drawfunctions

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/shreyghildiyal/goGame/camera"
	// "github.com/shreyghildiyal/goGame/spaceEntities"
)

func DrawGalaxy(screen *ebiten.Image, camera camera.Camera, systems map[int]int) {
	// DrawWarpLines(screen, camera, systems)
	// DrawGalaxyStars(screen, camera, systems)
}

func DrawWarpLines(screen *ebiten.Image, camera camera.Camera, systems map[int]map[int]int) {
	// colour := config.GetConfig().WarpLines.Colour

	// for _, system := range systems {
	// 	for _, system2Id := range system.Neighbours {
	// 		if system.Id < system2Id {
	// 			ebitenutil.DrawLine(screen, system.X+camera.X, system.Y+camera.Y, systems[system2Id].X+camera.X, systems[system2Id].Y+camera.Y, colour)
	// 		}
	// 	}

	// }
}

// func DrawGalaxyStars(screen *ebiten.Image, camera camera.Camera, stars map[int]*spaceEntities.System) {

// 	for _, system := range stars {

// 		DrawSpaceEntity(screen, camera.X, camera.Y, camera.Zoom, system)
// 	}
// }
