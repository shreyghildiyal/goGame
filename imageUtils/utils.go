package imageutils

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	config "github.com/shreyghildiyal/goGame/configs"
)

var (
	starImagesMap   map[string]*image.Image
	planetImagesMap map[string]*image.Image
)

func GetImage(imagePath string) image.Image {
	existingImageFile, err := os.Open(imagePath)
	if err != nil {
		// Handle error
		log.Fatal(err)
	}

	defer existingImageFile.Close()

	// Alternatively, since we know it is a png already
	// we can call png.Decode() directly
	loadedImage, err := png.Decode(existingImageFile)
	if err != nil {
		// Handle error
		log.Fatal(err)
	}
	// fmt.Println(loadedImage)
	return loadedImage
}

func GetSystemImage(systemImageType string) *ebiten.Image {
	fmt.Println("star image type", systemImageType)
	fmt.Println("Star images map is nil", starImagesMap == nil)
	return ebiten.NewImageFromImage(*starImagesMap[systemImageType])

}

func GetPlanetImage(planetImageType string) *ebiten.Image {
	return ebiten.NewImageFromImage(*planetImagesMap[planetImageType])
}

func InitImageMaps() {

	// initStarImageMap()
	starImagesMap = map[string]*image.Image{}
	initImageMap(starImagesMap, config.GetConfig().StarImages)
	planetImagesMap = map[string]*image.Image{}
	initImageMap(planetImagesMap, config.GetConfig().PlanetImages)
	// initPlanetImageMap()

	planetImagesMap = map[string]*image.Image{}
}

func initImageMap(imgmap map[string]*image.Image, pathMap map[string]string) {
	// imgmap = map[string]*image.Image{}

	for t, path := range pathMap {
		img := GetImage(path)
		imgmap[t] = &img
	}
}

func GetScale() (float64, float64) {
	return 1, 1
}
