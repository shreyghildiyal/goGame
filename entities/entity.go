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

func (e *Entity) AddComponent(componentId int, componentType constants.ComponentTypeName) error {

	if _, found := e.componentsMap[componentType]; !found {
		e.componentsMap[componentType] = []int{componentId}
		return nil
		// e.componentsMap[componentType] = append(e.componentsMap[componentType], )
	} else {

		for _, cId := range e.componentsMap[componentType] {
			if cId == componentId {
				return fmt.Errorf("component %d of type %s already added to entity %d", componentId, componentType, e.id)
			}
		}
		e.componentsMap[componentType] = append(e.componentsMap[componentType], componentId)
		return nil
	}

}
