package components

import (
	"fmt"

	"github.com/shreyghildiyal/goGame/constants"
	"github.com/shreyghildiyal/goGame/utils"
)

// type componentConstraint interface {
// 	Coordinates | Drawable | InSystem | Name
// }

type ComponentHandler[T Component] struct {
	componentsSlice []T
	freeIds         utils.Heap[int]
	freeIdsSlice    []bool
}

func (ch *ComponentHandler[T]) GetComponent(id int) (*T, error) {
	if id >= len(ch.componentsSlice) {
		return new(T), fmt.Errorf("no entity with id %d exists", id)
	} else if ch.freeIdsSlice[id] {
		found := false
		for _, freeId := range ch.freeIds.Dump() {
			if freeId == id {
				found = true
				break
			}
		}
		if !found {
			return new(T), fmt.Errorf("no entity at id %d. however the id is not available. something really wonky going on", id)
		}

		return new(T), fmt.Errorf("no entity with id %d exists", id)
	} else {
		return &ch.componentsSlice[id], nil
	}
}

func (ch *ComponentHandler[T]) AddComponent(entityId int, entityType constants.EntityTypeName) int {

	// component := T{markedForDelete: false, entityType: entityType, componentsMap: map[constants.ComponentTypeName][]int{}}
	// component := new(T)
	// (*component).MarkForDeletion()
	// *T.GetEntity()
	// component
	var component T
	component.SetEntity(entityId, entityType)
	componentId := -1
	if ch.freeIds.Len() == 0 {
		// component.setId(len(ch.componentsSlice))
		componentId = len(ch.componentsSlice)
		ch.componentsSlice = append(ch.componentsSlice, component)
		ch.freeIdsSlice = append(ch.freeIdsSlice, false)
	} else {
		enId := ch.freeIds.Pop()
		// component.id = enId
		componentId = enId
		ch.componentsSlice[enId] = component
		ch.freeIdsSlice[enId] = false
	}
	return componentId
}

func (ch *ComponentHandler[T]) DeleteComponent(componentId int) error {

	if componentId >= len(ch.componentsSlice) {
		return fmt.Errorf("no entity with id %d exists", componentId)
	} else if ch.freeIdsSlice[componentId] {
		return fmt.Errorf("no entity with id %d exists", componentId)
	} else {
		ch.componentsSlice[componentId] = *new(T)
		ch.freeIdsSlice[componentId] = true
		ch.freeIds.Push(componentId)
	}

	return nil
}
