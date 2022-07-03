package gamestate

import (
	"encoding/json"

	"github.com/shreyghildiyal/goGame/utils"
)

type SaveEntity struct {
	Id         int `json:"id"`
	EntityType int `json:"type"`
}

type SaveDrawable struct {
	TargetHeight float64 `json:"height"`
	TargetWidth  float64 `json:"width"`
	Image        string  `json:"imageName"`
	ZLevel       int     `json:"zLevel"`
	Id           int     `json:"id"`
	EntityId     int     `json:"entityId"`
}

type SaveCoordinates struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type SaveInSystem struct {
	SystemId    int             `json:"systemId"`
	Coordinates SaveCoordinates `json:"coordinates"`
	Id          int             `json:"id"`
	EntityId    int             `json:"entityId"`
}

type SaveComponents struct {
	SystemDrawables   []SaveDrawable `json:"systemDrawables"`
	GalacticDrawables []SaveDrawable `json:"galacticDrawables"`
	InSystem          []SaveInSystem `json:"inSystem"`
}

type SaveGame struct {
	Entities   []SaveEntity
	Components SaveComponents
}

func (gs *GameState) loadSaveGame() {

}

func (gs *GameState) saveGame() {

	saveObj := SaveGame{}

	for i := 0; i < gs.Entities.EntityListLen(); i++ {
		ent, err := gs.Entities.GetEntity(i)
		if err != nil {
			continue
		} else {
			saveEntity := SaveEntity{Id: ent.GetId()}
			saveObj.Entities = append(saveObj.Entities, saveEntity)
		}
	}

	jsonStr, _ := json.Marshal(saveObj)

	utils.SaveToFile(jsonStr, "quickSave.json")
}

// func loadEntities(entityContainer *entities.EntityHandler) {

// records, err := utils.GetCsvRecords(config.GetConfig().EntitiesFile)
// if err != nil {
// 	log.Fatal("Unable to read entities. Cant proceed", err)
// }

// for _, record := range records {
// 	entityId, err := strconv.Atoi(record[0])
// 	if err != nil {
// 		log.Fatal("Corrupted record", record, err)
// 	}
// 	entityTypeInt, err := strconv.Atoi(record[1])
// 	if err != nil {
// 		log.Fatal("Corrupted record", record, err)
// 	}
// 	entityType := constants.EntityTypeName(entityTypeInt)
// 	entityContainer.AddEntityWithId(entityId, entityType)
// }
// }

// func loadDrawables(drawable *components.ComponentHandler[*components.Drawable]) {

// records, err := utils.GetCsvRecords(config.GetConfig().DrawablesFile)
// if err != nil {
// 	log.Fatal("Unable to read entities. Cant proceed", err)
// }

// for _, record := range records {
// 	entityId, err := strconv.Atoi(record[0])
// 	if err != nil {
// 		log.Fatal("Corrupted record", record, err)
// 	}
// 	entityTypeInt, err := strconv.Atoi(record[1])
// 	if err != nil {
// 		log.Fatal("Corrupted record", record, err)
// 	}
// 	entityType := constants.EntityTypeName(entityTypeInt)
// 	drawable.AddEntityWithId(entityId, entityType)
// }
// }
