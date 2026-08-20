package drawing_test

import (
	"testing"

	"github.com/shreyghildiyal/goGame/drawing"
)

func TestAddGalaxyDrawable(t *testing.T) {

	dr := drawing.NewDrawableRegistry()

	draw := drawing.Drawable{}

	drawSlice := []drawing.Drawable{draw}

	dr.AddGalaxyDrawables("a", drawSlice)
}
