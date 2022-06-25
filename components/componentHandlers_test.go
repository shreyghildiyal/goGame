package components_test

import (
	"testing"

	"github.com/shreyghildiyal/goGame/components"
	"github.com/shreyghildiyal/goGame/constants"
	"github.com/shreyghildiyal/goGame/entities"
)

func TestComponentHandlerAddComponent(t *testing.T) {

	// ch := components.ComponentHandler[*components.Drawable]{}

	// componentId1 := ch.AddComponent(0, constants.STAR)
	// if componentId1 != 0 {
	// 	t.Errorf("The added component doesnt have the correct ID")
	// }

	// componentId2 := ch.AddComponent(0, constants.STAR)
	// if componentId2 != 1 {
	// 	t.Errorf("The added component doesnt have the correct ID")
	// }
}

func TestComponentHandlerGetComponent(t *testing.T) {

	// coord := components.Coordinates{}
	// coord.GetEnitityId()

	ch := components.ComponentHandler[*components.Drawable]{}

	drawable := components.Drawable{}
	entity := entities.Entity{}
	drawable.SetEntity(entity.GetId(), constants.EntityTypeName(entity.GetEntityType()))
	componentId1 := ch.AddComponent(&drawable)
	if componentId1 != 0 {
		t.Errorf("The added component doesnt have the correct ID")
	}

	component, err := ch.GetComponent(componentId1)
	if err != nil {
		t.Errorf("Error getting component from handler")
	}
	component.TargetHeight = 100
	component.TargetWidth = 200

	component2, err := ch.GetComponent(componentId1)
	if err != nil {
		t.Errorf("Error getting component from handler")
	}
	if component2.TargetHeight != component.TargetHeight || component2.TargetWidth != component.TargetWidth {
		t.Errorf("The component pointer is going to act difficult")
	}
}
