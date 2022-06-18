package drawfunctions

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/shreyghildiyal/goGame/camera"
	config "github.com/shreyghildiyal/goGame/configs"
	"github.com/shreyghildiyal/goGame/spaceEntities"
)

func DrawGalaxy(screen *ebiten.Image, camera camera.Camera, systems map[int]*spaceEntities.System) {
	DrawWarpLines(screen, camera, systems)
	DrawGalaxyStars(screen, camera, systems)
}

func DrawWarpLines(screen *ebiten.Image, camera camera.Camera, systems map[int]*spaceEntities.System) {
	colour := config.GetConfig().WarpLines.Colour

	for _, system := range systems {
		for _, system2Id := range system.Neighbours {
			if system.Id < system2Id {
				ebitenutil.DrawLine(screen, system.X+camera.X, system.Y+camera.Y, systems[system2Id].X+camera.X, systems[system2Id].Y+camera.Y, colour)
			}
		}

	}
}

func DrawGalaxyStars(screen *ebiten.Image, camera camera.Camera, stars map[int]*spaceEntities.System) {

	for _, system := range stars {

		DrawSpaceEntity(screen, camera.X, camera.Y, camera.Zoom, system)
	}
}
