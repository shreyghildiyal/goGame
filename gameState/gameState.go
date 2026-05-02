package gamestate

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/shreyghildiyal/goGame/camera"
	config "github.com/shreyghildiyal/goGame/configs"
	"github.com/shreyghildiyal/goGame/constants"
	drawfunctions "github.com/shreyghildiyal/goGame/drawFunctions"
	gameobjects "github.com/shreyghildiyal/goGame/gameobjects"
	imageutils "github.com/shreyghildiyal/goGame/imageUtils"
)

type ViewType string

const (
	MenuView   ViewType = "Menu"
	GalaxyView ViewType = "Galaxy"
	SystemView ViewType = "System"
)

// type View struct {
// 	ViewType ViewType
// 	Extra    interface{}
// }

// const (
// 	ScreenWidth  = 600
// 	ScreenHeight = 600
// )

type GameState struct {
	conf config.Configuration
	// Entities        entities.EntityHandler
	Background      *ebiten.Image
	CurrentSystemId int
	PrevUpdate      time.Time
	CurrentView     ViewType
	Camera          camera.Camera
	Keys            []ebiten.Key
	Galaxy          gameobjects.Galaxy

	// We want a list of all drawable entities. With Ids
	// We will then have a list of various base entities in each view.
	// For galaxy view, base entitites will be stars
	// For system view, base entities will be derived from the star object.
	// We probably also need an empires list.
	// in galaxyview, we will iterate through each star. we will find its corresponding drawable entity, and draw it
	// in system view, we will draw the star at the center. we will then get the planets and fleets in the star system.
	// then we will get the corresponding drawable entity and draw it
}

func (g *GameState) Update() error {

	dt := time.Since(g.PrevUpdate)

	HandleKeyboardInput(dt, g, g.conf.Camera)

	HandleMouseInput(dt, g)

	g.Keys = inpututil.AppendPressedKeys(g.Keys[:0])

	// for i := 0; i < len(g.Planets); i++ {
	// 	g.Planets[i].Update(dt)
	// }

	g.PrevUpdate = g.PrevUpdate.Add(dt)
	return nil
}

func (g *GameState) Draw(screen *ebiten.Image) {

	drawfunctions.DrawBackground(screen, g.Background)

	switch g.CurrentView {
	case GalaxyView:
		// drawfunctions.DrawGalaxy(screen, g.Camera, g.Systems)
	case SystemView:
		// fmt.Println("Drawing System")
		drawfunctions.DrawSystem(screen, &g.Camera, g.CurrentSystemId)
	case MenuView:
		drawfunctions.DrawMenu(screen)
	}

}

func Newgame(conf config.Configuration) *GameState {

	fmt.Println("planet", constants.PLANET)
	fmt.Println("star", constants.STAR)

	imageutils.InitImageMaps(conf)

	game := GameState{}
	game.conf = conf
	// game.Background = ebiten.NewImageFromImage(imageutils.GetImageFromMap(config.GetConfig().BackgroundImagePath))
	game.Background = imageutils.GetImageFromMap("backgroundImage")
	// game.Systems = spaceEntities.LoadSystems()
	// fmt.Println("Systems", game.Systems)
	// spaceEntities.CreateWarpLines(game.Systems)
	// game.Planets = spaceEntities.LoadPlanets()
	game.PrevUpdate = time.Now()
	game.CurrentView = GalaxyView
	game.Camera.Zoom = 1
	game.CurrentSystemId = 0

	game.loadSaveGame()

	// fmt.Println("Number of inSystems", game.inSystemHandler.Len())
	return &game
}

func (g *GameState) Layout(outsideWidth, outsideHeight int) (int, int) {

	return g.conf.ScreenSize.Width, g.conf.ScreenSize.Height
}
