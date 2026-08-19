package gamestate

func (gs *GameState) saveGame() {

	// jsonStr, err := json.Marshal(
	// 	struct {
	// 		Entities   *entities.EntityHandler `json:"entities"`
	// 		Components DummyComponents         `json:"components"`
	// 	}{
	// 		// Entities: &gs.Entities,
	// 		Components: DummyComponents{
	// 			// SystemDrawables: gs.systemDrawableHandler,
	// 			// GalaxyDrawables: gs.galaxyDrawableHandler,
	// 		},
	// 	},
	// )

	// if err == nil {
	// 	err = utils.SaveToFile(jsonStr, "quickSave.json", gs.conf.SaveGameDir)

	// }

	// if err != nil {
	// 	log.Println("Error saving game", err)
	// }

}
