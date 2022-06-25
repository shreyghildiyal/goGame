package entityManagement

import (
	"github.com/shreyghildiyal/goGame/components"
	"github.com/shreyghildiyal/goGame/constants"
	"github.com/shreyghildiyal/goGame/entities"
)

func AddPlanet(entityHandler *entities.EntityHandler, coordinateHandler *components.ComponentHandler[*components.Coordinates]) {
	// entity
	entityId := entityHandler.AddEntity(constants.PLANET)

	coordinates := components.Coordinates{}

	coordinates.SetEntity(entityId, constants.PLANET)

	coordinateHandler.AddComponent(&coordinates)

}
