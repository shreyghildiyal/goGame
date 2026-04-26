package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	config "github.com/shreyghildiyal/goGame/configs"
	gamestate "github.com/shreyghildiyal/goGame/gameState"
	"github.com/shreyghildiyal/goGame/gametext"
	imageutils "github.com/shreyghildiyal/goGame/imageUtils"
)

func main() {

	conf := config.GetConfig()
	imageutils.InitImageMaps(conf)
	gametext.InitFonts(conf)
	ebiten.SetWindowSize(conf.ScreenSize.Width, conf.ScreenSize.Height)
	ebiten.SetWindowTitle("Ebiten game practice")

	game := gamestate.Newgame(conf)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
