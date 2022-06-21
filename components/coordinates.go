package components

import "github.com/shreyghildiyal/goGame/constants"

type Coordinates struct {
	X            float64
	Y            float64
	entityId     int
	entityType   constants.EntityTypeName
	shouldDelete bool
}

func (c *Coordinates) GetEnitityId() int {
	return c.entityId
}

func (c *Coordinates) AsFloatPair() (float64, float64) {
	return c.X, c.Y
}

func (c *Coordinates) GetEntityType() constants.EntityTypeName {
	return c.entityType
}

func (c *Coordinates) SetEntity(entityId int, entityType constants.EntityTypeName) {
	c.entityId = entityId
	c.entityType = entityType
}

func (c *Coordinates) MarkForDeletion() {
	c.shouldDelete = true
}

func (c *Coordinates) IsMarkedForDelete() bool {
	return c.shouldDelete
}
