package drawing_test

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/shreyghildiyal/goGame/camera"
	"github.com/shreyghildiyal/goGame/drawing"
)

type TestImage struct {
}

func TestGetGeometry(t *testing.T) {

	d := drawing.Drawable{
		X:            0,
		Y:            0,
		Image:        ebiten.NewImage(100, 100),
		TargetHeight: 50,
		TargetWidth:  50,
	}

	cam := camera.Camera{
		Zoom: 0.5,
		X:    10,
		Y:    10,
	}

	geom := d.GetGeometry(200, 200, cam)

	x, y := geom.Apply(10, 10)

	if x != 215 || y != 215 {
		t.Error("The geometry seems to be be incorrect. got ", x, y)
	}
}
