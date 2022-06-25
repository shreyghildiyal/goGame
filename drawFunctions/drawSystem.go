package drawfunctions

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/shreyghildiyal/goGame/camera"
	"github.com/shreyghildiyal/goGame/components"
	"github.com/shreyghildiyal/goGame/constants"
	"github.com/shreyghildiyal/goGame/entities"
	// "github.com/shreyghildiyal/goGame/spaceEntities"
)

func DrawSystem(screen *ebiten.Image, camera *camera.Camera, currentSystemId int, entityList []entities.Entity, drawables []components.Drawable, inSystems []components.InSystem, coordinates []components.Coordinates) {
	// for _, planetId := range systems[systemId].Planets {
	// 	DrawPlanet(screen, planets[planetId])
	// }
	for _, drawable := range drawables {
		entityId := drawable.GetEntityId()
		inSystemIds, found := entityList[entityId].GetComponentIds(constants.INSYSTEM)
		if found {
			for _, inSysId := range inSystemIds {
				if inSysId == currentSystemId {
					// if systemCoordinates[]
					DrawSprite(screen, camera, drawable, components.Coordinates{X: 0, Y: 0})
				}
			}
		}
	}

	DrawSystemStar(screen, system)
}

func DrawSystemStar(screen *ebiten.Image, system *spaceEntities.System) {

}

// func DrawPlanet(screen *ebiten.Image, p *spaceEntities.Planet) {
// 	// fmt.Println("Planet Name", p.Name)
// 	x, y := p.Image.Size()
// 	op := &ebiten.DrawImageOptions{}
// 	op.GeoM.Translate(-float64(x/2), -float64(y/2))
// 	op.GeoM.Scale(p.Scale, p.Scale)
// 	op.GeoM.Translate(float64(x/2), float64(y/2))
// 	op.GeoM.Translate(p.X, p.Y)

// 	textX, textY := p.GetTextPosition(gametext.SpaceDisplayFont)

// 	if p.Image == nil {
// 		fmt.Println("image from", p.ImageLoc, "was nil somehow")
// 	} else if gametext.SpaceDisplayFont == nil {
// 		fmt.Println("font is nil")
// 	} else {
// 		screen.DrawImage(p.Image, op)
// 		// fmt.Printf("text location %d, %d\n", int(p.X), int(p.Y)+y+mplusNormalFont.Metrics().Height.Ceil())
// 		text.Draw(screen, p.Name, gametext.SpaceDisplayFont, textX, textY, config.GetConfig().Text.Colour)
// 	}

// }
