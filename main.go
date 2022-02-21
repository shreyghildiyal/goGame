package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	imageutils "github.com/shreyghildiyal/goGame/imageUtils"
	"github.com/shreyghildiyal/goGame/planets"
)

const (
	screenWidth  = 600
	screenHeight = 400

	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32
	frameNum    = 8
)

type Game struct {
	planets []planets.Planet
	// background *ebiten.Image
	count      int
	prevUpdate time.Time
}

func (g *Game) Update() error {

	dt := time.Since(g.prevUpdate)
	// fmt.Println("delta time", dt.Milliseconds())
	for i := 0; i < len(g.planets); i++ {
		g.planets[i].Update(dt)
	}

	g.prevUpdate.Add(dt)
	g.count++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	for _, p := range g.planets {
		p.Draw(screen)
	}
	// g.planets[0].Draw(screen)
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
	game := Game{}

	// game.background = ebiten.NewImageFromImage(imageutils.GetImage(config.GetConfig().BackgroundImagePath))
	game.planets = planets.LoadPlanets()
	game.prevUpdate = time.Now()
	fmt.Println("Number of planets", len(game.planets))
	return &game
}

func main() {
	// Decode an image from the image file's byte slice.
	// Now the byte slice is generated with //go:generate for Go 1.15 or older.
	// If you use Go 1.16 or newer, it is strongly recommended to use //go:embed to embed the image file.
	// See https://pkg.go.dev/embed for more details.
	// img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// runnerImage = ebiten.NewImageFromImage(img)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Animation (Ebiten Demo)")

	game := Newgame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func NewGame() {
	panic("unimplemented")
}
