package spaceEntities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	config "github.com/shreyghildiyal/goGame/configs"
	imageutils "github.com/shreyghildiyal/goGame/imageUtils"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const degreePerSec = math.Pi / 180000

var mplusNormalFont font.Face

const dpi = 72

type Planet struct {
	Id           int     `json:"id"`
	X            float64 `json:"x"`
	Y            float64 `json:"y"`
	image        *ebiten.Image
	ImageLoc     string  `json:"imageLoc"`
	Name         string  `json:"name"`
	Scale        float64 `json:"scale"`
	rotation     float64
	RotationRate float64 `json:"rotationRate"` // this is in degrees per second
	// nameFont     font.Face
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
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(float64(x/2), float64(y/2))
	op.GeoM.Translate(p.X, p.Y)
	op.GeoM.Scale(p.Scale, p.Scale)

	textX, textY := p.GetTextPosition(mplusNormalFont)

	if p.image == nil {
		fmt.Println("image from", p.ImageLoc, "was nil somehow")
	} else if mplusNormalFont == nil {
		fmt.Println("font is nil")
	} else {
		screen.DrawImage(p.image, op)
		// fmt.Printf("text location %d, %d\n", int(p.X), int(p.Y)+y+mplusNormalFont.Metrics().Height.Ceil())
		text.Draw(screen, p.Name, mplusNormalFont, textX, textY, config.GetConfig().Text.Colour)
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
	SetFont()
	// printTextPrintLoc(planets)
	planetMap := make(map[int]*Planet, len(planets))

	for _, planet := range planets {
		planetMap[planet.Id] = &planet
	}

	return planetMap
}

func (p *Planet) GetTextPosition(font font.Face) (int, int) {
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
	return int(p.X), int(p.Y)
}

func SetFont() {

	tt, _ := opentype.Parse(fonts.MPlus1pRegular_ttf)
	mplusNormalFont, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(config.GetConfig().Text.Size),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}
