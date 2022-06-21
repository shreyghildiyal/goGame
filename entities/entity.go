package entities

import (
	"github.com/shreyghildiyal/goGame/constants"
)

type Entity interface {
	GetId() int
	GetComponentIds(constants.ComponentTypeName) ([]int, bool)
	GetComponentsMap() map[constants.ComponentTypeName][]int
	SetId(int)
	AddComponent(constants.ComponentTypeName, int)
	GetEntityType() constants.EntityTypeName
}
