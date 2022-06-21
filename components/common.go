package components

import "github.com/shreyghildiyal/goGame/constants"

type Component interface {
	GetEntityId() int
	GetEntityType() constants.EntityTypeName
	SetEntity(int, constants.EntityTypeName)
	MarkForDeletion()
	IsMarkedForDelete() bool
}
