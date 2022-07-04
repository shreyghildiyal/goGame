package gamestate

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/shreyghildiyal/goGame/camera"
	"github.com/shreyghildiyal/goGame/components"
	config "github.com/shreyghildiyal/goGame/configs"
	"github.com/shreyghildiyal/goGame/constants"
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
	Entities        entities.EntityHandler
	Background      *ebiten.Image
	CurrentSystemId int
	PrevUpdate      time.Time
	CurrentView     ViewType
	Camera          camera.Camera
	Keys            []ebiten.Key

	// drawableHandler components.ComponentHandler[*components.Drawable]
	systemDrawableHandler components.ComponentHandler[*components.SystemDrawable]
	galaxyDrawableHandler components.ComponentHandler[*components.GalaxyDrawable]
	// coordinateHandler components.ComponentHandler[*components.Coordinates]
	inSystemHandler components.ComponentHandler[*components.InSystem]

	// systemCoordinateHandler components.ComponentHandler[*components.SystemCoordinates]
	// galaxyCoordinateHandler components.ComponentHandler[*components.GalaxyCoordinates]
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
		drawfunctions.DrawSystem(screen, &g.Camera, g.CurrentSystemId, g.Entities, g.systemDrawableHandler, g.inSystemHandler)
	case MenuView:
		drawfunctions.DrawMenu(screen)
	}

}

func Newgame() *GameState {

	fmt.Println("planet", constants.PLANET)
	fmt.Println("star", constants.STAR)

	imageutils.InitImageMaps()

	game := GameState{}

	// game.Background = ebiten.NewImageFromImage(imageutils.GetImageFromMap(config.GetConfig().BackgroundImagePath))
	game.Background = imageutils.GetImageFromMap("backgroundImage")
	// game.Systems = spaceEntities.LoadSystems()
	// fmt.Println("Systems", game.Systems)
	// spaceEntities.CreateWarpLines(game.Systems)
	// game.Planets = spaceEntities.LoadPlanets()
	game.PrevUpdate = time.Now()
	game.CurrentView = SystemView
	game.Camera.Zoom = 1
	game.CurrentSystemId = 0

	game.Entities.AddEntity(constants.STAR)
	game.Entities.AddEntity(constants.STAR)
	game.Entities.AddEntity(constants.STAR)
	// game.Entities.AddEntityWithId(6, constants.STAR)
	game.Entities.DeleteEntity(1)

	game.loadSaveGame()

	fmt.Println("Number of inSystems", game.inSystemHandler.Len())
	return &game
}

func (g *GameState) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.GetConfig().ScreenSize.Width, config.GetConfig().ScreenSize.Height
}
