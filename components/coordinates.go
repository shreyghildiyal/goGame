package components

type Coordinates struct {
	X float64
	Y float64
}

func (c *Coordinates) AsFloatPair() (float64, float64) {
	return c.X, c.Y
}
