package entityManagement

import (
	"github.com/shreyghildiyal/goGame/components"
	"github.com/shreyghildiyal/goGame/entities"
)

func AddPlanet(entityHandler *entities.EntityHandler, drawables []components.Drawable, inSystem []components.InSystem, coords []components.Coordinates) {
	// entity
	planet := entities.Planet{}
	// coords := components.Coordinates{}
	entityHandler.AddEntity(&planet)

}
