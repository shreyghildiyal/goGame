package gamestate

import (
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

type View struct {
	ViewType ViewType
	Extra    interface{}
}

// const (
// 	ScreenWidth  = 600
// 	ScreenHeight = 600
// )

type Game struct {
	// Planets         map[int]*spaceEntities.Planet
	// Systems         map[int]*spaceEntities.System
	Entities        []entities.Entity
	Background      *ebiten.Image
	CurrentSystemId int
	PrevUpdate      time.Time

	CurrentView View

	CurrentSystemID int

	Camera camera.Camera

	Keys []ebiten.Key

	Drawables []components.Drawable

	InSystemComponents []components.InSystem

	CoordinateComponents []components.Coordinates
}

func (g *Game) Update() error {

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

func (g *Game) Draw(screen *ebiten.Image) {

	drawfunctions.DrawBackground(screen, g.Background)

	switch g.CurrentView.ViewType {
	case GalaxyView:
		// drawfunctions.DrawGalaxy(screen, g.Camera, g.Systems)
	case SystemView:
		drawfunctions.DrawSystem(screen, &g.Camera, g.CurrentSystemId, g.Entities, g.Drawables, g.InSystemComponents, g.CoordinateComponents)
	case MenuView:
		drawfunctions.DrawMenu(screen)
	}

}

func Newgame() *Game {

	imageutils.InitImageMaps()

	game := Game{}

	game.Background = ebiten.NewImageFromImage(imageutils.GetImage(config.GetConfig().BackgroundImagePath))
	// game.Systems = spaceEntities.LoadSystems()
	// fmt.Println("Systems", game.Systems)
	// spaceEntities.CreateWarpLines(game.Systems)
	// game.Planets = spaceEntities.LoadPlanets()
	game.PrevUpdate = time.Now()
	game.CurrentView = View{ViewType: GalaxyView}
	game.Camera.Zoom = 1
	// fmt.Println("Number of planets", len(game.Planets))
	return &game
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.GetConfig().ScreenSize.Width, config.GetConfig().ScreenSize.Height
}
