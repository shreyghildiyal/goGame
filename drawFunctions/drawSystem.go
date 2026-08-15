package drawfunctions

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/shreyghildiyal/goGame/camera"
	"github.com/shreyghildiyal/goGame/gameobjects"
	// "github.com/shreyghildiyal/goGame/spaceEntities"
)

func DrawSystem(
	screen *ebiten.Image,
	camera *camera.Camera,
	currentSystemId int,
	g gameobjects.Galaxy,

) error {
	star, found := g.Stars[currentSystemId]

	if !found {
		return fmt.Errorf("No star found with given ID")
	} else {
		star.SystemSprite.Draw(screen, *camera)
	}
	return nil
}

// func DrawSystemStar(screen *ebiten.Image, system *spaceEntities.System) {

// }

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
