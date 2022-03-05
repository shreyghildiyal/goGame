package spaceEntities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	config "github.com/shreyghildiyal/goGame/configs"
	imageutils "github.com/shreyghildiyal/goGame/imageUtils"
	"golang.org/x/image/font"
)

type System struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Planets  []Planet `json:"planets"`
	X        float64  `json:"x"`
	Y        float64  `json:"y"`
	Scale    float64  `json:"scale"`
	StarType string   `json:"starType"`
	image    *ebiten.Image
}

func LoadSystems() map[int]*System {
	systemsFile := config.GetConfig().SystemsFile

	data, err := ioutil.ReadFile(systemsFile)

	if err != nil {
		log.Fatal(err)
	}

	systems := []System{}

	err = json.Unmarshal(data, &systems)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(systems); i++ {

		systems[i].image = imageutils.GetSystemImage(systems[i].StarType)

	}

	systemsMap := make(map[int]*System, len(systems))

	for _, system := range systems {
		systemsMap[system.Id] = &system
	}

	return systemsMap
}

func DrawWarpLines(screen *ebiten.Image, x, y, z int, stars map[int]*System) {

}

func DrawStars(screen *ebiten.Image, x, y, z int, stars map[int]*System) {

	for _, system := range stars {
		system.Draw(screen, x, y, z)
	}
}

func (star *System) Draw(screen *ebiten.Image, camX, camYy, camZoom int) {

	x, y := star.image.Size()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(x/2), -float64(y/2))
	// op.GeoM.Rotate(star.rotation)
	op.GeoM.Translate(float64(x/2), float64(y/2))
	op.GeoM.Translate(star.X, star.Y)
	op.GeoM.Scale(star.Scale, star.Scale)

	textX, textY := star.GetTextPosition(mplusNormalFont)

	if star.image == nil {
		fmt.Println("image from", star.StarType, "was nil somehow")
	} else if mplusNormalFont == nil {
		fmt.Println("font is nil")
	} else {
		screen.DrawImage(star.image, op)
		// fmt.Printf("text location %d, %d\n", int(p.X), int(p.Y)+y+mplusNormalFont.Metrics().Height.Ceil())
		text.Draw(screen, star.Name, mplusNormalFont, textX, textY, config.GetConfig().Text.Colour)
	}
}

func (system *System) GetTextPosition(font font.Face) (int, int) {
	// _, y := p.image.Size()
	// textX := int(p.X)

	// textY := 0.0
	// // fmt.Println("textHeight1", textY)
	// textY += p.Y
	// fmt.Println("textHeight2", textY, p.Y)
	// textY += p.Scale * float64(y)
	// fmt.Println("textHeight3", textY, y)
	// textY += float64(font.Metrics().CapHeight.Ceil())
	// fmt.Println("textHeight4", textY, font.Metrics().CapHeight.Ceil())

	// return textX, int(textY)
	return int(system.X), int(system.Y)
}
