package drawfunctions

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/shreyghildiyal/goGame/camera"
	"github.com/shreyghildiyal/goGame/gameobjects"
	// "github.com/shreyghildiyal/goGame/spaceEntities"
)

func DrawGalaxy(screen *ebiten.Image, camera camera.Camera, galaxy gameobjects.Galaxy) {

	for _, star := range galaxy.Stars {
		star.Draw(screen, camera)
		// _ = star
	}
}
