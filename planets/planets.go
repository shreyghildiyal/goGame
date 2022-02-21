package planets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	config "github.com/shreyghildiyal/goGame/configs"
	imageutils "github.com/shreyghildiyal/goGame/imageUtils"
)

const rotationDamper = 0.00001

type Planet struct {
	X            float64 `json:"x"`
	Y            float64 `json:"y"`
	image        *ebiten.Image
	ImageLoc     string  `json:"imageLoc"`
	Name         string  `json:"name"`
	Scale        float64 `json:"scale"`
	rotation     float64
	RotationRate float64 `json:"rotationRate"`
}

func (p *Planet) Update(dt time.Duration) {

	p.rotation += rotationDamper * float64(dt.Milliseconds()) * p.RotationRate
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

	if p.image == nil {
		fmt.Println("image from", p.ImageLoc, "was nil somehow")
	} else {
		screen.DrawImage(p.image, op)
	}

}

func LoadPlanets() []Planet {
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

	return planets
}
