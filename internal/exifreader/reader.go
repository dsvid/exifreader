package exifreader

import (
	"os"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
)

func ExtractCameraModel(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	exifData, err := exif.Decode(f)
	if err != nil {
		return "", err
	}

	model, err := exifData.Get(exif.Model)
	if err != nil {
		return "", err
	}

	return strings.Trim(model.String(), "\""), nil
}
