package spaceEntities

import (
	imageutils "github.com/shreyghildiyal/goGame/imageUtils"
)

type SpaceEntity interface {
	GetDisplay() *imageutils.DispDetails
	GetCoordinates() (float64, float64)
	// GetTextPosition(font font.Face) (int, int)
	GetName() string
}

func IsClicked(entity SpaceEntity, cursorX, cursorY float64, camX, camY, camZoom float64) bool {
	disp := entity.GetDisplay()
	dX := float64(disp.TargetWidth / 2)
	dY := float64(disp.TargetHeight / 2)
	x, y := entity.GetCoordinates()
	if x-dX <= cursorX-camX && cursorX-camX <= x+dX && y-dY <= cursorY-camY && cursorY-camY <= y+dY {
		return true
	}
	return false
}
