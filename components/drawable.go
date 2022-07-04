package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/shreyghildiyal/goGame/constants"
	imageutils "github.com/shreyghildiyal/goGame/imageUtils"
)

type Drawable struct {
	TargetHeight float64 `json:"baseHeight"`
	TargetWidth  float64 `json:"baseWidth"`
	ScaleX       float64 `json:"scaleX"`
	ScaleY       float64 `json:"scaleY"`
	ZLevel       int     `json:"zLevel"`
	ImageName    string  `json:"imageName"`
	Image        *ebiten.Image
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
		ImageName:     imageName,
		baseComponent: baseComponent{entityId: entityId, entityType: entityType},
	}

	// drawable.ScaleX = drawable.TargetWidth / float64(drawable.Image.Bounds().Dx())
	// drawable.ScaleY = drawable.TargetHeight / float64(drawable.Image.Bounds().Dy())

	return drawable
}

// func (d *Drawable) MarshalJSON([]byte, error) {
// 	fmt.Println("MarshalJSON called for entity handler")
// 	j, err := json.Marshal(struct {
// 		ComponentsSlice []T
// 		FreeIds         utils.NumberHeap[int]
// 		FreeIdsSlice    []bool
// 	}{
// 		ComponentsSlice: ch.componentsSlice,
// 		FreeIds:         ch.freeIds,
// 		FreeIdsSlice:    ch.freeIdsSlice,
// 	})
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}
// 	return j, nil
// }
