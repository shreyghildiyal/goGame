package imageutils

import (
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
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

// Register image decoders so image.Decode knows how to parse them
func init() {
	image.RegisterFormat("jpeg", "\xff\xd8", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "\x89PNG\r\n\x1a\n", png.Decode, png.DecodeConfig)
}

func getImage(imagePath string) image.Image {
	existingImageFile, err := os.Open(imagePath)
	if err != nil {
		// Handle error
		log.Fatal(err)
	}

	defer existingImageFile.Close()

	// Alternatively, since we know it is a png already
	// we can call png.Decode() directly

	loadedImage, _, err := image.Decode(existingImageFile)

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

func InitImageMaps(conf config.Configuration) error {

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
	imgPathMap, err := imagePathMap(conf.ImagesFile)
	if err != nil {
		fmt.Println("imgPathMap loading error")
		return err
	}
	initImageMap(imagesMap, imgPathMap)
	fmt.Println("Number of images in map", len(imagesMap))
	return nil
}

func imagePathMap(imagesFile string) (map[string]string, error) {
	imgFileMap := map[string]string{}

	res, err := os.ReadFile(imagesFile)
	if err != nil {
		fmt.Println("failed to read file", imagesFile)
		return nil, err
	}
	err = json.Unmarshal(res, &imgFileMap)
	if err != nil {
		fmt.Println("failed to unmarshal json")
		return nil, err
	}

	return imgFileMap, nil
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
