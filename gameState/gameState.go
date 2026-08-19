package gamestate

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/shreyghildiyal/goGame/camera"
	config "github.com/shreyghildiyal/goGame/configs"
	drawfunctions "github.com/shreyghildiyal/goGame/drawFunctions"
	"github.com/shreyghildiyal/goGame/gameobjects"
	imageutils "github.com/shreyghildiyal/goGame/imageUtils"
)

type ViewType string

const (
	MenuView   ViewType = "Menu"
	GalaxyView ViewType = "Galaxy"
	SystemView ViewType = "System"
)

type GameState struct {
	conf config.Configuration

	Background      *ebiten.Image
	CurrentSystemId int
	PrevUpdate      time.Time
	CurrentView     ViewType
	Camera          camera.Camera
	Keys            []ebiten.Key
	Galaxy          gameobjects.Galaxy
}

func (g *GameState) Update() error {

	dt := time.Since(g.PrevUpdate)

	HandleKeyboardInput(dt, g, g.conf.Camera)

	HandleMouseInput(dt, g)

	g.Keys = inpututil.AppendPressedKeys(g.Keys[:0])

	g.PrevUpdate = g.PrevUpdate.Add(dt)
	return nil
}

func (g *GameState) Draw(screen *ebiten.Image) {

	drawfunctions.DrawBackground(screen, g.Background)

	switch g.CurrentView {
	case GalaxyView:
		err := drawfunctions.DrawGalaxy(screen, g.Camera, g.Galaxy)
		if err != nil {
			log.Fatal(err)
		}
	case SystemView:
		// fmt.Println("Drawing System")
		err := drawfunctions.DrawSystem(screen, &g.Camera, g.CurrentSystemId, g.Galaxy)
		if err != nil {
			log.Fatal(err)
		}
	case MenuView:
		drawfunctions.DrawMenu(screen)
	}

}

func Newgame(conf config.Configuration) (*GameState, error) {

	err := imageutils.InitImageMaps(conf)
	if err != nil {
		return nil, err
	}

	game := GameState{}
	game.Galaxy, err = gameobjects.LoadGalaxy(conf.SaveGameFile, conf)
	if err != nil {
		return nil, err
	}

	game.conf = conf

	game.Background, err = imageutils.GetImageFromMap("backgroundImage")
	if err != nil {
		return nil, err
	}

	game.PrevUpdate = time.Now()
	game.CurrentView = GalaxyView
	game.Camera.Zoom = 1
	game.Camera.X = 0
	game.Camera.Y = 0
	game.CurrentSystemId = 0
	game.CurrentView = GalaxyView

	// game.loadSaveGame()

	// fmt.Println("Number of inSystems", game.inSystemHandler.Len())
	return &game, nil
}

func (g *GameState) Layout(outsideWidth, outsideHeight int) (int, int) {

	return g.conf.ScreenSize.Width, g.conf.ScreenSize.Height
}
