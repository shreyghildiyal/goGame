package gametext

import (
	"bytes"
	"fmt"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
	config "github.com/shreyghildiyal/goGame/configs"
)

var (
	// Global face reference matching your legacy pattern
	SpaceDisplayFont *text.GoTextFace
	SpaceColour      color.Color
)

func loadFontBytes(fontPath string) ([]byte, error) {
	// Read the raw file bytes into a []byte slice
	fontBytes, err := os.ReadFile(fontPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read font file from %s: %w", fontPath, err)
	}

	return fontBytes, nil
}

// InitFonts accepts dynamic font bytes (or file paths) and target font size.
func InitFonts(conf config.TextConf) error {

	fontBytes, err := loadFontBytes(conf.FontFile)
	if err != nil {
		return err
	}

	// Parse OpenType/TrueType source
	source, err := text.NewGoTextFaceSource(bytes.NewReader(fontBytes))
	if err != nil {
		return fmt.Errorf("failed to parse font: %w", err)
	}

	// Create face instance
	SpaceDisplayFont = &text.GoTextFace{
		Source: source,
		Size:   conf.Size,
	}

	return nil
}
