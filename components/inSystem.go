package components

import "github.com/shreyghildiyal/goGame/constants"

type InSystem struct {
	systemId     int
	entityId     int
	entityType   constants.EntityTypeName
	shouldDelete bool
}

func (is *InSystem) GetEntityId() int {
	return is.entityId
}

func (is *InSystem) GetEntityType() constants.EntityTypeName {
	return is.entityType
}

func (is *InSystem) SetEntity(entityId int, entityType constants.EntityTypeName) {
	is.entityId = entityId
	is.entityType = entityType
}

func (is *InSystem) SetSystemId(systemId int) {
	is.systemId = systemId
}

func (is *InSystem) GetSystemId() int {
	return is.systemId
}

func (is *InSystem) MarkForDeletion() {
	is.shouldDelete = true
}

func (is *InSystem) IsMarkedForDelete() bool {
	return is.shouldDelete
}
