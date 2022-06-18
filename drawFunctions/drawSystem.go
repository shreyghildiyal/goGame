package drawfunctions

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	config "github.com/shreyghildiyal/goGame/configs"
	"github.com/shreyghildiyal/goGame/gametext"
	"github.com/shreyghildiyal/goGame/spaceEntities"
)

func DrawSystem(screen *ebiten.Image, system *spaceEntities.System, planets map[int]*spaceEntities.Planet) {
	for _, planetId := range system.Planets {
		DrawPlanet(screen, planets[planetId])
	}

	DrawSystemStar(screen, system)
}

func DrawSystemStar(screen *ebiten.Image, system *spaceEntities.System) {

}

func DrawPlanet(screen *ebiten.Image, p *spaceEntities.Planet) {
	// fmt.Println("Planet Name", p.Name)
	x, y := p.Image.Size()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(x/2), -float64(y/2))
	op.GeoM.Scale(p.Scale, p.Scale)
	op.GeoM.Translate(float64(x/2), float64(y/2))
	op.GeoM.Translate(p.X, p.Y)

	textX, textY := p.GetTextPosition(gametext.SpaceDisplayFont)

	if p.Image == nil {
		fmt.Println("image from", p.ImageLoc, "was nil somehow")
	} else if gametext.SpaceDisplayFont == nil {
		fmt.Println("font is nil")
	} else {
		screen.DrawImage(p.Image, op)
		// fmt.Printf("text location %d, %d\n", int(p.X), int(p.Y)+y+mplusNormalFont.Metrics().Height.Ceil())
		text.Draw(screen, p.Name, gametext.SpaceDisplayFont, textX, textY, config.GetConfig().Text.Colour)
	}

}
