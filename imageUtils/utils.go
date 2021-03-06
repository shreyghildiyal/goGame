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
	// starImagesMap   map[string]*image.Image
	// planetImagesMap map[string]*image.Image
	imagesMap map[string]*image.Image
)

func getImage(imagePath string) image.Image {
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

func GetImageFromMap(imageName string) *ebiten.Image {

	// fmt.Println("ImagesMap is nil", imagesMap == nil)
	// if imagesMap == nil {
	// 	InitImageMaps()
	// }
	fmt.Println("Image Name", imageName)
	if _, found := imagesMap[imageName]; !found {
		fmt.Println("Image Name", imageName, "not present in map")
	}

	return ebiten.NewImageFromImage(*imagesMap[imageName])
}

// func GetSystemImage(systemImageType string) *ebiten.Image {
// 	fmt.Println("star image type", systemImageType)
// 	// fmt.Println("Star images map is nil", starImagesMap == nil)
// 	return ebiten.NewImageFromImage(*starImagesMap[systemImageType])
// }

// func GetPlanetImage(planetImageType string) *ebiten.Image {
// 	fmt.Println("planetImagesMap is nil", planetImagesMap == nil)
// 	for key, val := range planetImagesMap {
// 		fmt.Println(key, ":", val)
// 	}
// 	return ebiten.NewImageFromImage(*planetImagesMap[planetImageType])
// }

func InitImageMaps() {

	fmt.Println("Initializing image maps")
	// initStarImageMap()
	// starImagesMap = map[string]*image.Image{}
	// initImageMap(starImagesMap, config.GetConfig().StarImages)
	// planetImagesMap = map[string]*image.Image{}
	// fmt.Println("Loading planet images using", config.GetConfig().PlanetImages)

	// initImageMap(planetImagesMap, config.GetConfig().PlanetImages)

	// fmt.Println("Number of planet images in map", len(planetImagesMap))
	// initPlanetImageMap()
	imagesMap = map[string]*image.Image{}
	initImageMap(imagesMap, config.GetConfig().Images)
	fmt.Println("Number of images in map", len(imagesMap))
}

func initImageMap(imgmap map[string]*image.Image, pathMap map[string]string) {
	// imgmap = map[string]*image.Image{}

	for t, path := range pathMap {
		img := getImage(path)
		imgmap[t] = &img
	}
}

func GetScale() (float64, float64) {
	return 1, 1
}
