package spaceEntities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	config "github.com/shreyghildiyal/goGame/configs"
	"github.com/shreyghildiyal/goGame/gametext"
	imageutils "github.com/shreyghildiyal/goGame/imageUtils"
	"golang.org/x/image/font"
)

// var mplusNormalFont font.Face

type Planet struct {
	Id       int     `json:"id"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	image    *ebiten.Image
	ImageLoc string  `json:"imageLoc"`
	Name     string  `json:"name"`
	Scale    float64 `json:"scale"`
}

func (p *Planet) Update(dt time.Duration) {
	// fmt.Println("Delta time", dt.Milliseconds())
	// p.rotation += degreePerSec * float64(dt.Milliseconds()) * p.RotationRate
}

func (p *Planet) Draw(screen *ebiten.Image) {
	// fmt.Println("Planet Name", p.Name)
	x, y := p.image.Size()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(x/2), -float64(y/2))
	op.GeoM.Scale(p.Scale, p.Scale)
	op.GeoM.Translate(float64(x/2), float64(y/2))
	op.GeoM.Translate(p.X, p.Y)

	textX, textY := p.GetTextPosition(gametext.SpaceDisplayFont)

	if p.image == nil {
		fmt.Println("image from", p.ImageLoc, "was nil somehow")
	} else if gametext.SpaceDisplayFont == nil {
		fmt.Println("font is nil")
	} else {
		screen.DrawImage(p.image, op)
		// fmt.Printf("text location %d, %d\n", int(p.X), int(p.Y)+y+mplusNormalFont.Metrics().Height.Ceil())
		text.Draw(screen, p.Name, gametext.SpaceDisplayFont, textX, textY, config.GetConfig().Text.Colour)
	}

}

func LoadPlanets() map[int]*Planet {
	planetsFilePath := config.GetConfig().PlanetsFile

	data, err := ioutil.ReadFile(planetsFilePath)

	if err != nil {
		log.Fatal(err)
	}

	planets := []Planet{}

	err = json.Unmarshal(data, &planets)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(planets); i++ {

		planets[i].image = ebiten.NewImageFromImage(imageutils.GetImage(planets[i].ImageLoc))

	}

	// printTextPrintLoc(planets)
	planetMap := make(map[int]*Planet, len(planets))

	for i := 0; i < len(planets); i++ {
		planetMap[planets[i].Id] = &planets[i]
	}

	return planetMap
}

func (p *Planet) GetTextPosition(font font.Face) (int, int) {
	return int(p.X), int(p.Y)
}
