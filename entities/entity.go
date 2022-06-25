package entities

import (
	"fmt"

	"github.com/shreyghildiyal/goGame/constants"
)

type Entity struct {
	id              int
	componentsMap   map[constants.ComponentTypeName][]int
	entityType      constants.EntityTypeName
	markedForDelete bool
}

func (e *Entity) GetId() int {
	return e.id
}

func (e *Entity) GetComponentIds(componentType constants.ComponentTypeName) ([]int, error) {
	if compIds, found := e.componentsMap[componentType]; found {
		return compIds, nil
	} else {
		return []int{}, fmt.Errorf("no components of the given type associated with type %s", componentType)
	}
}

func (e *Entity) GetEntityType() constants.EntityTypeName {
	return e.entityType
}

func (e *Entity) MarkForDeletion() {
	e.markedForDelete = true
}

func (e *Entity) IsMarkedForDeletion() bool {
	return e.markedForDelete
}
