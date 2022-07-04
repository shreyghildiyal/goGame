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

type SystemDrawable Drawable
type GalaxyDrawable Drawable

func NewDrawable[T SystemDrawable | GalaxyDrawable](entityId int, entityType constants.EntityType, imageName string, height, width float64) T {

	image := imageutils.GetImageFromMap(imageName)
	scaleX := width / float64(image.Bounds().Dx())
	scaleY := height / float64(image.Bounds().Dy())
	drawable := T{
		Image:         imageutils.GetImageFromMap(imageName),
		TargetHeight:  height,
		TargetWidth:   width,
		ScaleX:        scaleX,
		ScaleY:        scaleY,
		baseComponent: baseComponent{entityId: entityId, entityType: entityType},
	}

	// drawable.ScaleX = drawable.TargetWidth / float64(drawable.Image.Bounds().Dx())
	// drawable.ScaleY = drawable.TargetHeight / float64(drawable.Image.Bounds().Dy())

	return drawable
}
