package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/shreyghildiyal/goGame/constants"
	imageutils "github.com/shreyghildiyal/goGame/imageUtils"
)

type Drawable struct {
	TargetHeight float64 `json:"baseHeight"`
	TargetWidth  float64 `json:"baseWidth"`
	ScaleX       float64
	ScaleY       float64
	Image        *ebiten.Image
	ZLevel       int
	baseComponent
}

func NewDrawable(entityId int, entityType constants.EntityType, imageName string, height, width float64) Drawable {
	drawable := Drawable{
		Image:         imageutils.GetPlanetImage(imageName),
		TargetHeight:  height,
		TargetWidth:   width,
		baseComponent: baseComponent{entityId: entityId, entityType: entityType},
	}

	drawable.ScaleX = drawable.TargetWidth / float64(drawable.Image.Bounds().Dx())
	drawable.ScaleY = drawable.TargetHeight / float64(drawable.Image.Bounds().Dy())

	return drawable
}
