package components

import "github.com/shreyghildiyal/goGame/constants"

type Component interface {
	GetEntityId() int
	GetEntityType() constants.EntityTypeName
	SetEntity(int, constants.EntityTypeName)
	MarkForDeletion()
	IsMarkedForDelete() bool
	GetId() int
	SetId(int)
}

type baseComponent struct {
	id           int
	entityId     int
	entityType   constants.EntityTypeName
	shouldDelete bool
}

func (is *baseComponent) GetEntityId() int {
	return is.entityId
}

func (c *baseComponent) GetEntityType() constants.EntityTypeName {
	return c.entityType
}

func (c *baseComponent) SetEntity(entityId int, entityType constants.EntityTypeName) {
	c.entityId = entityId
	c.entityType = entityType
}

func (c *baseComponent) MarkForDeletion() {
	c.shouldDelete = true
}

func (c *baseComponent) IsMarkedForDelete() bool {
	return c.shouldDelete
}

func (c *baseComponent) GetId() int {
	return c.id
}
func (c *baseComponent) SetId(id int) {
	c.id = id
}
