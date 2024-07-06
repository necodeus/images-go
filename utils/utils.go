package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
)

func GetExtension(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".png" || ext == ".jpg" || ext == ".jpeg" || ext == ".gif" || ext == ".bmp" || ext == ".tiff" || ext == ".webp"
}

func GenerateThumbnail(srcPath, dstPath string, width, height int) error {
	srcImage, err := imaging.Open(srcPath)
	if err != nil {
		return fmt.Errorf("failed to open image: %v", err)
	}

	scaledImage := imaging.Fit(srcImage, width, height, imaging.Lanczos)

	dstImage := imaging.CropCenter(scaledImage, width, height)

	err = imaging.Save(dstImage, dstPath)

	if err != nil {
		return fmt.Errorf("failed to save image: %v", err)
	}

	err = os.Rename(dstPath, strings.TrimSuffix(dstPath, filepath.Ext(dstPath)))

	if err != nil {
		return fmt.Errorf("failed to rename image: %v", err)
	}

	return nil
}

func ParseResolution(resolution string) (int, int, error) {
	parts := strings.Split(resolution, "x")

	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid resolution format")
	}

	width, err := strconv.Atoi(parts[0])

	if err != nil {
		return 0, 0, fmt.Errorf("invalid width")
	}

	height, err := strconv.Atoi(parts[1])

	if err != nil {
		return 0, 0, fmt.Errorf("invalid height")
	}

	return width, height, nil
}

func IsValidResolution(resolutions []string, width, height int) bool {
	for _, res := range resolutions {
		w, h, err := ParseResolution(res)

		if err == nil && w == width && h == height {
			return true
		}
	}

	return false
}
