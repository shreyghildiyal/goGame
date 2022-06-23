package entities

import (
	"fmt"

	"github.com/shreyghildiyal/goGame/constants"
	"github.com/shreyghildiyal/goGame/utils"
)

type Entity interface {
	GetId() int
	GetComponentIds(constants.ComponentTypeName) ([]int, bool)
	GetComponentsMap() map[constants.ComponentTypeName][]int
	SetId(int)
	AddComponent(constants.ComponentTypeName, int)
	GetEntityType() constants.EntityTypeName
	MarkForDelete()
	IsMarkedForDelete() bool
}

type EntityHandler struct {
	entities []Entity
	freeIds  utils.Heap[int]
}

func (eh *EntityHandler) GetEntity(id int) (Entity, error) {
	if id >= len(eh.entities) {
		return nil, fmt.Errorf("no entity with id %d exists", id)
	} else if eh.entities[id] == nil {
		found := false
		for _, freeId := range eh.freeIds.Dump() {
			if freeId == id {
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("no entity at id %d. however the id is not available. something really wonky going on", id)
		}

		return nil, fmt.Errorf("no entity with id %d exists", id)
	} else {
		return eh.entities[id], nil
	}
}

func (eh *EntityHandler) AddEntity(entity Entity) {

	if eh.freeIds.Len() == 0 {
		entity.SetId(eh.freeIds.Len())
		eh.entities = append(eh.entities, entity)
	} else {
		enId := eh.freeIds.Pop()
		entity.SetId(enId)
		eh.entities = append(eh.entities, entity)
	}

}

func (eh *EntityHandler) DeleteEntity(entityId int) error {

	if entityId >= len(eh.entities) {
		return fmt.Errorf("no entity with id %d exists", entityId)
	} else if eh.entities[entityId] == nil {
		return fmt.Errorf("no entity with id %d exists", entityId)
	} else {
		eh.entities[entityId] = nil
		eh.freeIds.Push(entityId)
	}

	return nil
}
