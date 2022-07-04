package entities

import (
	"encoding/json"
	"fmt"

	"github.com/shreyghildiyal/goGame/constants"
	"github.com/shreyghildiyal/goGame/utils"
)

type EntityHandler struct {
	entities     []Entity
	freeIds      utils.NumberHeap[int]
	freeIdsSlice []bool
}

func (eh *EntityHandler) GetEntity(id int) (*Entity, error) {
	if id >= len(eh.entities) {
		return &Entity{id: -1}, fmt.Errorf("no entity with id %d exists", id)
	} else if eh.freeIdsSlice[id] {
		found := false
		for _, freeId := range eh.freeIds.Dump() {
			if freeId == id {
				found = true
				break
			}
		}
		if !found {
			return &Entity{id: -1}, fmt.Errorf("no entity at id %d. however the id is not available. something really wonky going on", id)
		}

		return &Entity{id: -1}, fmt.Errorf("no entity with id %d exists", id)
	} else {
		return &eh.entities[id], nil
	}
}

func (eh *EntityHandler) AddEntity(entityType constants.EntityType) int {

	entity := Entity{markedForDelete: false, entityType: entityType, componentsMap: map[constants.ComponentTypeName][]int{}}
	if eh.freeIds.Len() == 0 {
		entity.id = len(eh.entities)
		eh.entities = append(eh.entities, entity)
		eh.freeIdsSlice = append(eh.freeIdsSlice, false)
	} else {
		enId := eh.freeIds.Pop()
		entity.id = enId
		eh.entities[enId] = entity
		eh.freeIdsSlice[enId] = false
	}
	return entity.id
}

func (eh *EntityHandler) DeleteEntity(entityId int) error {

	if entityId >= len(eh.entities) {
		return fmt.Errorf("no entity with id %d exists", entityId)
	} else if eh.freeIdsSlice[entityId] {
		return fmt.Errorf("no entity with id %d exists", entityId)
	} else {
		eh.entities[entityId] = Entity{id: -1}
		eh.freeIdsSlice[entityId] = true
		eh.freeIds.Push(entityId)
	}

	return nil
}

func (eh *EntityHandler) AddComponentToEntity(entityId int, componentId int, componentType constants.ComponentTypeName) error {
	if entityId >= len(eh.entities) {
		return fmt.Errorf("no entity with id %d exists", entityId)
	} else if eh.freeIdsSlice[entityId] {
		return fmt.Errorf("no entity with id %d exists", entityId)
	} else {
		eh.entities[entityId].AddComponent(componentId, componentType)
	}
	return nil
}

func (eh *EntityHandler) AddEntityWithId(entityId int, entityType constants.EntityType) error {
	return nil
}

func (eh *EntityHandler) EntityListLen() int {
	return len(eh.entities)
}

func (eh *EntityHandler) MarshalJSON() ([]byte, error) {
	fmt.Println("MarshalJSON called for entity handler")
	j, err := json.Marshal(struct {
		Entities     []Entity              `json:"entities"`
		FreeIds      utils.NumberHeap[int] `json:"freeIds"`
		FreeIdsSlice []bool                `json:"freeIdsSlice"`
	}{
		Entities:     eh.entities,
		FreeIds:      eh.freeIds,
		FreeIdsSlice: eh.freeIdsSlice,
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return j, nil
}
