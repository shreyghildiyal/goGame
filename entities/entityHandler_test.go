package entities_test

import (
	"testing"

	"github.com/shreyghildiyal/goGame/constants"
	"github.com/shreyghildiyal/goGame/entities"
)

func TestEntityHandler(t *testing.T) {
	entityHandler := entities.EntityHandler{}

	enId := entityHandler.AddEntity(constants.PLANET)

	if enId != 0 {
		t.Error("The Id of a new entity in an empty handler should be 0")
	}

	enId2 := entityHandler.AddEntity(constants.STAR)

	if enId2 != 1 {
		t.Error("The Id of new entity is wrong")
	}

	ent1, err := entityHandler.GetEntity(0)
	if err != nil {
		t.Error("Error getting objhect from handler")
	}
	if ent1.GetId() != 0 {
		t.Errorf("Wrong id for the returned entity. Expected %d found %d", 0, ent1.GetId())
	}

	_, err = entityHandler.GetEntity(3)
	if err == nil {
		t.Error("The entity should have not been found")
	}

	_ = entityHandler.AddEntity(constants.STAR)
	_ = entityHandler.AddEntity(constants.STAR)

	err = entityHandler.DeleteEntity(10)
	if err == nil {
		t.Error("We cant have deleted an entity that doesnt exist")
	}

	err = entityHandler.DeleteEntity(1)
	if err != nil {
		t.Error("We should be able to delete an entity")
	}

	enId = entityHandler.AddEntity(constants.STAR)
	if enId != 1 {
		t.Errorf("new entity should have the free id 1. but has %d", enId)
	}
}
