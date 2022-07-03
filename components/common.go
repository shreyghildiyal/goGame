package components

import "github.com/shreyghildiyal/goGame/constants"

type Component interface {
	GetEntityId() int
	GetEntityType() constants.EntityType
	SetEntity(int, constants.EntityType)
	MarkForDeletion()
	IsMarkedForDelete() bool
	GetId() int
	SetId(int)
}

type baseComponent struct {
	id           int
	entityId     int
	entityType   constants.EntityType
	shouldDelete bool
}

func (is *baseComponent) GetEntityId() int {
	return is.entityId
}

func (c *baseComponent) GetEntityType() constants.EntityType {
	return c.entityType
}

func (c *baseComponent) SetEntity(entityId int, entityType constants.EntityType) {
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
