package entityManagement

import (
	"fmt"

	"github.com/shreyghildiyal/goGame/components"
	"github.com/shreyghildiyal/goGame/constants"
	"github.com/shreyghildiyal/goGame/entities"
	imageutils "github.com/shreyghildiyal/goGame/imageUtils"
)

func AddPlanet(entityHandler *entities.EntityHandler, coordinateHandler *components.ComponentHandler[*components.Coordinates], drawableHandler *components.ComponentHandler[*components.Drawable]) {
	// entity
	entityId := entityHandler.AddEntity(constants.PLANET)

	coordinates := components.Coordinates{X: 100, Y: 100}
	drawable := components.Drawable{Image: imageutils.GetPlanetImage("planetType1")}

	coordinates.SetEntity(entityId, constants.PLANET)
	coordinateHandler.AddComponent(&coordinates)
	fmt.Println("Added coordinates to planet")
	fmt.Println(coordinateHandler.Len())

	drawable.SetEntity(entityId, constants.PLANET)
	drawableHandler.AddComponent(&drawable)
}
