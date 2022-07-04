package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	config "github.com/shreyghildiyal/goGame/configs"
	gamestate "github.com/shreyghildiyal/goGame/gameState"
	"github.com/shreyghildiyal/goGame/gametext"
	imageutils "github.com/shreyghildiyal/goGame/imageUtils"
)

// func GetImages(paths []string) []*ebiten.Image {
// 	images := make([]*ebiten.Image, len(paths))

// 	for i, path := range paths {
// 		images[i] = ebiten.NewImageFromImage(imageutils.GetImage(path))
// 	}
// 	return images
// }

func main() {

	imageutils.InitImageMaps()
	gametext.InitFonts()
	ebiten.SetWindowSize(config.GetConfig().ScreenSize.Width, config.GetConfig().ScreenSize.Height)
	ebiten.SetWindowTitle("Ebiten game practice")

	game := gamestate.Newgame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
