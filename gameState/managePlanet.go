package gamestate

import (
	"github.com/shreyghildiyal/goGame/components"
	"github.com/shreyghildiyal/goGame/constants"
	"github.com/shreyghildiyal/goGame/entities"
)

func AddPlanet(entityHandler *entities.EntityHandler, gs *GameState) {
	// entity
	entityId := entityHandler.AddEntity(constants.PLANET)

	systemCoordinates := components.NewCoordinates[components.SystemCoordinates](200, 200)
	systemCoordinates.SetEntity(entityId, constants.PLANET)
	gs.systemCoordinateHandler.AddComponent(&systemCoordinates)

	// drawable := components.Drawable{Image: imageutils.GetPlanetImage("planetType1")}
	// drawable.SetEntity(entityId, constants.PLANET)
	drawable := components.NewDrawable(entityId, constants.PLANET, "planetType1", 50, 50)
	gs.drawableHandler.AddComponent(&drawable)
	gs.Entities.AddComponentToEntity(entityId, drawable.GetId(), constants.INSYSTEM)

	inSystem := components.NewInSystemComponent(0, entityId, constants.PLANET)
	gs.inSystemHandler.AddComponent(&inSystem)
	gs.Entities.AddComponentToEntity(entityId, inSystem.GetId(), constants.INSYSTEM)
}
