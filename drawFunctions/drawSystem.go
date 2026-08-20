package drawfunctions

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/shreyghildiyal/goGame/camera"
	"github.com/shreyghildiyal/goGame/drawing"
	"github.com/shreyghildiyal/goGame/gameobjects"
)

func DrawSystem(
	screen *ebiten.Image,
	camera *camera.Camera,
	currentSystemId string,
	g gameobjects.Galaxy,
	drawReg drawing.DrawableRegistry,
) error {
	star, found := g.Stars[currentSystemId]

	if !found {
		return fmt.Errorf("No star found with given ID %s", currentSystemId)
	} else {
		drawables, err := drawReg.GetSystemDrawables(star.Id)
		if err != nil {
			return err
		}
		for _, drawable := range drawables {
			drawable.Draw(screen, *camera)
		}

		for _, planetId := range star.Planets {
			drawables, err := drawReg.GetSystemDrawables(planetId)
			if err != nil {
				return err
			}
			for _, drawable := range drawables {
				drawable.Draw(screen, *camera)
			}
			DrawCircle(screen, *camera, g.Planets[planetId].SysRad)
		}
		// star.SystemSprite.Draw(screen, *camera)
		DrawCircle(screen, *camera, star.SystemRadius)
	}
	return nil
}

func DrawCircle(screen *ebiten.Image, camera camera.Camera, radius float64) {
	cx, cy := camera.WorldToScreen(0, 0, float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy()))
	scaledRadius := float32(radius * camera.Zoom)
	// fmt.Println("drawing circle", "cx", cx, "cy", cy, "rad", scaledRadius)
	vector.StrokeCircle(screen, cx, cy, scaledRadius, 1, color.RGBA{0, 150, 255, 100}, true)
}
