package gamestate

import (
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

	// g.Camera.X -= 5 * dt.Seconds()
	// g.Camera.Y -= 4 * dt.Seconds()

	g.PrevUpdate = g.PrevUpdate.Add(dt)
	return nil
}

func (g *GameState) Draw(screen *ebiten.Image) {

	drawfunctions.DrawBackground(screen, g.Background)

	switch g.CurrentView {
	case GalaxyView:
		drawfunctions.DrawGalaxy(screen, g.Camera, g.Galaxy)
	case SystemView:
		// fmt.Println("Drawing System")
		drawfunctions.DrawSystem(screen, &g.Camera, g.CurrentSystemId, g.Galaxy)
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
	// game.Galaxy = gameobjects.Galaxy{
	// 	Stars: map[int]gameobjects.Star{
	// 		1: gameobjects.NewStar(
	// 			1,
	// 			"one",
	// 			drawing.Drawable{
	// 				X:            250,
	// 				Y:            250,
	// 				TargetHeight: 20,
	// 				TargetWidth:  20,
	// 				Image:        imageutils.GetImageFromMap("redDwarf"),
	// 				RotAngle:     0,
	// 			},
	// 			drawing.Drawable{
	// 				X:            0,
	// 				Y:            0,
	// 				TargetHeight: 100,
	// 				TargetWidth:  100,
	// 				Image:        imageutils.GetImageFromMap("redDwarf"),
	// 				RotAngle:     0,
	// 			},
	// 		),
	// 		2: gameobjects.NewStar(
	// 			2,
	// 			"two",
	// 			drawing.Drawable{
	// 				X:            150,
	// 				Y:            250,
	// 				TargetHeight: 30,
	// 				TargetWidth:  30,
	// 				Image:        imageutils.GetImageFromMap("redDwarf"),
	// 				RotAngle:     0,
	// 			},
	// 			drawing.Drawable{
	// 				X:            0,
	// 				Y:            0,
	// 				TargetHeight: 100,
	// 				TargetWidth:  100,
	// 				Image:        imageutils.GetImageFromMap("redDwarf"),
	// 				RotAngle:     0,
	// 			},
	// 		),
	// 		3: gameobjects.NewStar(
	// 			3,
	// 			"tree",
	// 			drawing.Drawable{
	// 				X:            50,
	// 				Y:            250,
	// 				TargetHeight: 10,
	// 				TargetWidth:  10,
	// 				Image:        imageutils.GetImageFromMap("redDwarf"),
	// 				RotAngle:     0,
	// 			},
	// 			drawing.Drawable{
	// 				X:            0,
	// 				Y:            0,
	// 				TargetHeight: 100,
	// 				TargetWidth:  100,
	// 				Image:        imageutils.GetImageFromMap("redDwarf"),
	// 				RotAngle:     0,
	// 			},
	// 		),
	// 	},
	// }
	game.conf = conf
	// game.Background = ebiten.NewImageFromImage(imageutils.GetImageFromMap(config.GetConfig().BackgroundImagePath))
	game.Background, err = imageutils.GetImageFromMap("backgroundImage")
	if err != nil {
		return nil, err
	}
	// game.Systems = spaceEntities.LoadSystems()
	// fmt.Println("Systems", game.Systems)
	// spaceEntities.CreateWarpLines(game.Systems)
	// game.Planets = spaceEntities.LoadPlanets()
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
