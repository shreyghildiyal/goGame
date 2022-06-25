package components_test

import (
	"testing"

	"github.com/shreyghildiyal/goGame/components"
	"github.com/shreyghildiyal/goGame/constants"
)

func TestComponentHandlerAddComponent(t *testing.T) {

	ch := components.ComponentHandler[*components.Drawable]{}

	componentId1 := ch.AddComponent(0, constants.STAR)
	if componentId1 != 0 {
		t.Errorf("The added component doesnt have the correct ID")
	}

	componentId2 := ch.AddComponent(0, constants.STAR)
	if componentId2 != 1 {
		t.Errorf("The added component doesnt have the correct ID")
	}
}

func TestComponentHandlerGetComponent(t *testing.T) {

	// coord := components.Coordinates{}
	// coord.GetEnitityId()

	ch := components.ComponentHandler[*components.Drawable]{}

	componentId1 := ch.AddComponent(0, constants.PLANET)
	if componentId1 != 0 {
		t.Errorf("The added component doesnt have the correct ID")
	}

	component, err := ch.GetComponent(componentId1)
	if err != nil {
		t.Errorf("Error getting component from handler")
	}
	(*component).TargetHeight = 100
	(*component).TargetWidth = 200

	component2, err := ch.GetComponent(componentId1)
	if err != nil {
		t.Errorf("Error getting component from handler")
	}
	if (*component2).TargetHeight != (*component).TargetHeight || component2.Y != component2.Y {
		t.Errorf("The component pointer is going to act difficult")
	}
}
