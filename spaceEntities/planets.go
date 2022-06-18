package spaceEntities

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	config "github.com/shreyghildiyal/goGame/configs"
	imageutils "github.com/shreyghildiyal/goGame/imageUtils"
	"golang.org/x/image/font"
)

// var mplusNormalFont font.Face

type Planet struct {
	Id       int     `json:"id"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Image    *ebiten.Image
	ImageLoc string  `json:"imageLoc"`
	Name     string  `json:"name"`
	Scale    float64 `json:"scale"`
}

func (p *Planet) Update(dt time.Duration) {
	// fmt.Println("Delta time", dt.Milliseconds())
	// p.rotation += degreePerSec * float64(dt.Milliseconds()) * p.RotationRate
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

		planets[i].Image = ebiten.NewImageFromImage(imageutils.GetImage(planets[i].ImageLoc))

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
