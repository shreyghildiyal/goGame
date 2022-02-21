package imageutils

import (
	"image"
	"image/png"
	"log"
	"os"
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
