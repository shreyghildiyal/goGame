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
	err := imageutils.InitImageMaps(conf)
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = gametext.InitFonts(conf.Text)
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(conf.ScreenSize.Width, conf.ScreenSize.Height)
	ebiten.SetWindowTitle("Ebiten game practice")

	game, err := gamestate.Newgame(conf)

	if err != nil {
		log.Fatalln(err)
		return
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
