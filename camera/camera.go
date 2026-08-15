package camera

type Camera struct {
	X    float64
	Y    float64
	Zoom float64
}

func (c *Camera) ResetGalaxyView(x float64, y float64) {
	c.X = x
	c.Y = y
	c.Zoom = 1
}

func (c *Camera) ResetSystemDefault() {
	c.X = 0
	c.Y = 0
	c.Zoom = 1
}
