package drawfunctions

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/shreyghildiyal/goGame/camera"
	"github.com/shreyghildiyal/goGame/drawing"
	"github.com/shreyghildiyal/goGame/gameobjects"
	"github.com/shreyghildiyal/goGame/gametext"
)

func DrawGalaxy(screen *ebiten.Image, camera camera.Camera, galaxy gameobjects.Galaxy, drawReg drawing.DrawableRegistry) error {

	for _, star := range galaxy.Stars {
		// star.DrawGalaxySprite(screen, camera)

		drawables, err := drawReg.GetGalaxyDrawables(star.Id)
		if err != nil {
			return err
		}
		for _, drawable := range drawables {
			drawable.Draw(screen, camera)
			drawable.DrawName(star.Name, gametext.SpaceDisplayFont, gametext.SpaceColour, screen, camera)
		}

	}
	return nil
}
