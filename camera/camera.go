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

func (c *Camera) WorldToScreen(worldX, worldY float64, screenW, screenH float64) (float32, float32) {
	screenX := (screenW / 2) + (worldX-c.X)*c.Zoom
	screenY := (screenH / 2) + (worldY-c.Y)*c.Zoom
	return float32(screenX), float32(screenY)
}
