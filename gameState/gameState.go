package gamestate

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/shreyghildiyal/goGame/camera"
	"github.com/shreyghildiyal/goGame/components"
	config "github.com/shreyghildiyal/goGame/configs"
	drawfunctions "github.com/shreyghildiyal/goGame/drawFunctions"
	"github.com/shreyghildiyal/goGame/entities"
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
	// Planets         map[int]*spaceEntities.Planet
	// Systems         map[int]*spaceEntities.System
	Entities        entities.EntityHandler
	Background      *ebiten.Image
	CurrentSystemId int
	PrevUpdate      time.Time
	CurrentView     ViewType
	Camera          camera.Camera
	Keys            []ebiten.Key

	drawableHandler components.ComponentHandler[*components.Drawable]
	// coordinateHandler components.ComponentHandler[*components.Coordinates]
	inSystemHandler components.ComponentHandler[*components.InSystem]

	systemCoordinateHandler components.ComponentHandler[*components.SystemCoordinates]
	galaxyCoordinateHandler components.ComponentHandler[*components.GalaxyCoordinates]
}

func (g *GameState) Update() error {

	dt := time.Since(g.PrevUpdate)

	HandleKeyboardInput(dt, g)

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
		drawfunctions.DrawSystem(screen, &g.Camera, g.CurrentSystemId, g.Entities, g.drawableHandler, g.inSystemHandler, g.systemCoordinateHandler)
	case MenuView:
		drawfunctions.DrawMenu(screen)
	}

}

func Newgame() *GameState {

	imageutils.InitImageMaps()

	game := GameState{}

	game.Background = ebiten.NewImageFromImage(imageutils.GetImage(config.GetConfig().BackgroundImagePath))
	// game.Systems = spaceEntities.LoadSystems()
	// fmt.Println("Systems", game.Systems)
	// spaceEntities.CreateWarpLines(game.Systems)
	// game.Planets = spaceEntities.LoadPlanets()
	game.PrevUpdate = time.Now()
	game.CurrentView = SystemView
	game.Camera.Zoom = 1
	game.CurrentSystemId = 0
	// fmt.Println("Number of planets", len(game.Planets))

	AddPlanet(&game.Entities, &game)

	fmt.Println("Number of systemCoordinates", game.systemCoordinateHandler.Len())
	return &game
}

func (g *GameState) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.GetConfig().ScreenSize.Width, config.GetConfig().ScreenSize.Height
}
