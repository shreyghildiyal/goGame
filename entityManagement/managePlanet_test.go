package entityManagement_test

import (
	"testing"

	"github.com/shreyghildiyal/goGame/components"
	"github.com/shreyghildiyal/goGame/entities"
	"github.com/shreyghildiyal/goGame/entityManagement"
)

func TestSomething(t *testing.T) {
	eh := entities.EntityHandler{}
	ch := components.ComponentHandler[*components.Coordinates]{}
	dh := components.ComponentHandler[*components.Drawable]{}
	entityManagement.AddPlanet(&eh, &ch, &dh)
}
