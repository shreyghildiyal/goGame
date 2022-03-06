package spaceEntities

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	config "github.com/shreyghildiyal/goGame/configs"
	imageutils "github.com/shreyghildiyal/goGame/imageUtils"
	"golang.org/x/image/font"
)

type System struct {
	Id         int                    `json:"id"`
	Name       string                 `json:"name"`
	Planets    []Planet               `json:"planets"`
	X          float64                `json:"x"`
	Y          float64                `json:"y"`
	Display    imageutils.DispDetails `json:"displayDetails"`
	StarType   string                 `json:"starType"`
	Neighbours []*System
	// height   int
	// width    int
	// image *ebiten.Image
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

		systems[i].Display.Image = imageutils.GetSystemImage(systems[i].StarType)
		systems[i].Display.InitDetails()
		// systems[i].Display.BaseScaleX, systems[i].Display.BaseScaleY = imageutils.GetScale(systems[i].Display.Image, systems[i].X, systems[i].Y)
	}

	systemsMap := make(map[int]*System, len(systems))

	for i := 0; i < len(systems); i++ {
		systemsMap[systems[i].Id] = &systems[i]
	}

	return systemsMap
}

func DrawWarpLines(screen *ebiten.Image, camX, camY, zoom float64, systems map[int]*System) {
	colour := config.GetConfig().WarpLines.Colour

	for _, system := range systems {
		for _, system2 := range system.Neighbours {
			if system.Id < system2.Id {
				ebitenutil.DrawLine(screen, system.X+camX+float64(system.Display.BaseWidth)/2, system.Y+camY+float64(system.Display.BaseHeight)/2, system2.X+camX+float64(system2.Display.BaseWidth)/2, system2.Y+camY+float64(system2.Display.BaseHeight)/2, colour)
			}
		}

	}
}

func (s *System) GetCoordinates() (float64, float64) {
	return s.X, s.Y
}

func (s *System) GetDisplay() *imageutils.DispDetails {
	return &s.Display
}

func (s *System) GetName() string {
	return s.Name
}

func DrawStars(screen *ebiten.Image, camX, camY, zoom float64, stars map[int]*System) {

	for _, system := range stars {

		DrawSpaceEntity(screen, camX, camY, zoom, system)
	}
}

func (system *System) GetTextPosition(font font.Face) (int, int) {

	return int(system.X), int(system.Y)
}

func CreateWarpLines(systems map[int]*System) {
	systemConnectionsFile := config.GetConfig().SystemsConnectionFile

	data, err := ioutil.ReadFile(systemConnectionsFile)

	if err != nil {
		log.Fatal(err)
	}

	out := [][]int{}

	err = json.Unmarshal(data, &out)
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range out {
		sys1, found := systems[line[0]]
		if !found {
			continue
		}
		sys2, found2 := systems[line[1]]
		if !found2 {
			continue
		}
		if sys1.Id < sys2.Id {
			sys1.Neighbours = append(sys1.Neighbours, sys2)
			sys2.Neighbours = append(sys2.Neighbours, sys1)
		}

	}
}
