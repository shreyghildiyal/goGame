package gamestate

import (
	"encoding/json"
	"log"

	"github.com/shreyghildiyal/goGame/components"
	"github.com/shreyghildiyal/goGame/entities"
	"github.com/shreyghildiyal/goGame/utils"
)

func (gs *GameState) loadSaveGame() {

}

type DummyComponents struct {
	SystemDrawables components.ComponentHandler[*components.SystemDrawable] `json:"systemDrawables"`
	GalaxyDrawables components.ComponentHandler[*components.GalaxyDrawable] `json:"galaxyDrawables"`
}

func (gs *GameState) saveGame() {

	jsonStr, err := json.Marshal(
		struct {
			Entities   *entities.EntityHandler `json:"entities"`
			Components DummyComponents         `json:"components"`
		}{
			// Entities: &gs.Entities,
			Components: DummyComponents{
				// SystemDrawables: gs.systemDrawableHandler,
				// GalaxyDrawables: gs.galaxyDrawableHandler,
			},
		},
	)

	if err == nil {
		err = utils.SaveToFile(jsonStr, "quickSave.json", gs.conf.SaveGameDir)

	}

	if err != nil {
		log.Println("Error saving game", err)
	}

}
