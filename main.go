package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	config "github.com/shreyghildiyal/goGame/configs"
	"github.com/shreyghildiyal/goGame/gametext"
	imageutils "github.com/shreyghildiyal/goGame/imageUtils"
	"github.com/shreyghildiyal/goGame/inputs"
	"github.com/shreyghildiyal/goGame/spaceEntities"
)

const (
	screenWidth  = 600
	screenHeight = 600
)

type ViewType string

const (
	MenuView   ViewType = "Menu"
	GalaxyView ViewType = "Galaxy"
	SystemView ViewType = "System"
)

type View struct {
	viewType ViewType
	extra    interface{}
}

type Camera struct {
	x    float64
	y    float64
	zoom float64
}

type Game struct {
	planets    map[int]*spaceEntities.Planet
	systems    map[int]*spaceEntities.System
	background *ebiten.Image
	// count      int
	prevUpdate time.Time

	mouse inputs.Mouse

	currentView *View

	currentSystem *spaceEntities.System

	camera Camera

	keys []ebiten.Key
}

func (g *Game) handleKeyboardInput(dt time.Duration) {
	var timeMulti float64 = float64(dt.Microseconds()) / 1000000
	// fmt.Println("time Multi", timeMulti)
	// fmt.Println("Speed", config.GetConfig().Camera.SpeedX, config.GetConfig().Camera.SpeedY)
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.camera.y += config.GetConfig().Camera.SpeedY * timeMulti
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.camera.y -= config.GetConfig().Camera.SpeedY * timeMulti
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.camera.x += config.GetConfig().Camera.SpeedX * timeMulti
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.camera.x -= config.GetConfig().Camera.SpeedX * timeMulti
	}
}

func (g *Game) handleMouseInput(dt time.Duration) {

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if g.currentView.viewType == GalaxyView {
			fmt.Println("Just licked left mouse button")
			for _, system := range g.systems {
				cursorX, cursorY := ebiten.CursorPosition()
				if spaceEntities.IsClicked(system, float64(cursorX), float64(cursorY), g.camera.x, g.camera.y, g.camera.zoom) {
					fmt.Println("System", system.Name, "was clicked")
					// set view to system view
					// set current system to the system in question
				}
			}
		}
	}
}

func (g *Game) Update() error {

	dt := time.Since(g.prevUpdate)

	g.handleKeyboardInput(dt)

	g.handleMouseInput(dt)

	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	for i := 0; i < len(g.planets); i++ {
		g.planets[i].Update(dt)
	}

	g.prevUpdate = g.prevUpdate.Add(dt)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	g.drawBackground(screen)

	switch g.currentView.viewType {
	case GalaxyView:
		g.drawGalaxy(screen)
	case SystemView:
		g.drawSystem(screen, g.currentSystem)
	case MenuView:
		g.drawMenu(screen)
	}

}

func (g *Game) drawGalaxy(screen *ebiten.Image) {
	spaceEntities.DrawWarpLines(screen, g.camera.x, g.camera.y, g.camera.zoom, g.systems)
	spaceEntities.DrawStars(screen, g.camera.x, g.camera.y, g.camera.zoom, g.systems)
}

func (g *Game) drawMenu(screen *ebiten.Image) {

}

func (g *Game) drawSystem(screen *ebiten.Image, system *spaceEntities.System) {
	for _, p := range system.Planets {
		p.Draw(screen)
	}
}

func (g *Game) drawBackground(screen *ebiten.Image) {
	screen.DrawImage(g.background, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func GetImages(paths []string) []*ebiten.Image {
	images := make([]*ebiten.Image, len(paths))

	for i, path := range paths {
		images[i] = ebiten.NewImageFromImage(imageutils.GetImage(path))
	}
	return images
}

func Newgame() *Game {

	imageutils.InitImageMaps()

	game := Game{}

	game.background = ebiten.NewImageFromImage(imageutils.GetImage(config.GetConfig().BackgroundImagePath))
	game.systems = spaceEntities.LoadSystems()
	fmt.Println("Systems", game.systems)
	spaceEntities.CreateWarpLines(game.systems)
	game.planets = spaceEntities.LoadPlanets()
	game.prevUpdate = time.Now()
	game.currentView = &View{viewType: GalaxyView}
	game.camera.zoom = 1
	fmt.Println("Number of planets", len(game.planets))
	return &game
}

func main() {

	imageutils.InitImageMaps()
	gametext.InitFonts()
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Animation (Ebiten Demo)")

	game := Newgame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
