package gamestate

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/shreyghildiyal/goGame/camera"
	config "github.com/shreyghildiyal/goGame/configs"
	drawfunctions "github.com/shreyghildiyal/goGame/drawFunctions"
	imageutils "github.com/shreyghildiyal/goGame/imageUtils"
	"github.com/shreyghildiyal/goGame/spaceEntities"
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

type Game struct {
	Planets    map[int]*spaceEntities.Planet
	Systems    map[int]*spaceEntities.System
	Background *ebiten.Image
	// count      int
	PrevUpdate time.Time

	// mouse inputs.Mouse

	CurrentView ViewType

	CurrentSystemID int

	Camera camera.Camera

	Keys []ebiten.Key
}

func (g *Game) Update() error {

	dt := time.Since(g.PrevUpdate)

	HandleKeyboardInput(dt, g)

	HandleMouseInput(dt, g)

	g.Keys = inpututil.AppendPressedKeys(g.Keys[:0])

	for i := 0; i < len(g.Planets); i++ {
		g.Planets[i].Update(dt)
	}

	g.PrevUpdate = g.PrevUpdate.Add(dt)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	drawfunctions.DrawBackground(screen, g.Background)

	switch g.CurrentView {
	case GalaxyView:
		drawfunctions.DrawGalaxy(screen, g.Camera, g.Systems)
	case SystemView:
		drawfunctions.DrawSystem(screen, g.Systems[g.CurrentSystemID], g.Planets)
	case MenuView:
		drawfunctions.DrawMenu(screen)
	}

}

func Newgame() *Game {

	imageutils.InitImageMaps()

	game := Game{}

	game.Background = ebiten.NewImageFromImage(imageutils.GetImage(config.GetConfig().BackgroundImagePath))
	game.Systems = spaceEntities.LoadSystems()
	fmt.Println("Systems", game.Systems)
	spaceEntities.CreateWarpLines(game.Systems)
	game.Planets = spaceEntities.LoadPlanets()
	game.PrevUpdate = time.Now()
	game.CurrentView = GalaxyView
	game.Camera.Zoom = 1
	fmt.Println("Number of planets", len(game.Planets))
	return &game
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.GetConfig().ScreenSize.Width, config.GetConfig().ScreenSize.Height
}
