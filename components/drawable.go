package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/shreyghildiyal/goGame/constants"
)

type Drawable struct {
	TargetHeight float64 `json:"baseHeight"`
	TargetWidth  float64 `json:"baseWidth"`
	ScaleX       float64
	ScaleY       float64
	Image        *ebiten.Image
	ZLevel       int
	entityId     int
	entityType   constants.EntityTypeName
	shouldDelete bool
}

func (d *Drawable) GetEntityId() int {
	return d.entityId
}

func (d *Drawable) GetEntityType() constants.EntityTypeName {
	return d.entityType
}

func (d *Drawable) SetEntity(entityId int, entityType constants.EntityTypeName) {
	d.entityId = entityId
	d.entityType = entityType
}

func (d *Drawable) MarkForDeletion() {
	d.shouldDelete = true
}

func (d *Drawable) IsMarkedForDelete() bool {
	return d.shouldDelete
}
