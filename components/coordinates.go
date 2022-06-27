package components

type Coordinates struct {
	X float64
	Y float64
	baseComponent
}

func (c *Coordinates) AsFloatPair() (float64, float64) {
	return c.X, c.Y
}

type SystemCoordinates Coordinates

type GalaxyCoordinates Coordinates

func NewCoordinates[T SystemCoordinates | GalaxyCoordinates](x, y float64) T {
	return T{X: x, Y: y}
}
