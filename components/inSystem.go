package components

import "github.com/shreyghildiyal/goGame/constants"

type InSystem struct {
	systemId int
	Coordinates
	baseComponent
}

func (is *InSystem) SetSystemId(systemId int) {
	is.systemId = systemId
}

func (is *InSystem) GetSystemId() int {
	return is.systemId
}

func NewInSystemComponent(systemId int, entityId int, entityType constants.EntityType, x, y float64) InSystem {
	return InSystem{
		systemId: systemId,
		baseComponent: baseComponent{
			entityId:   entityId,
			entityType: entityType,
		},
		Coordinates: Coordinates{X: x, Y: y}}
}
