package constants

type ComponentTypeName string

const (
	COORDINATES ComponentTypeName = "COORDINATES"
	MOVABLE     ComponentTypeName = "MOVABLE"
	DRAWABLE    ComponentTypeName = "DRAWABLE"
	INSYSTEM    ComponentTypeName = "INSYSTEM"
	NAME        ComponentTypeName = "NAME"
)

type EntityTypeName string

const (
	PLANET EntityTypeName = "PLANET"
	STAR   EntityTypeName = "STAR"
)
