package constants

type ComponentTypeName string

const (
	COORDINATES ComponentTypeName = "COORDINATES"
	MOVABLE     ComponentTypeName = "MOVABLE"
	DRAWABLE    ComponentTypeName = "DRAWABLE"
	INSYSTEM    ComponentTypeName = "INSYSTEM"
	NAME        ComponentTypeName = "NAME"
)

type EntityType int

const (
	PLANET EntityType = iota
	STAR   EntityType = iota
)
