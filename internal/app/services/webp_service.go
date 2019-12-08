package services

import (
	"context"
	"github.com/davidbyttow/govips/pkg/vips"
	"os"
)

func Convert(ctx context.Context, image *os.File) (*os.File, error) {
	file, err := os.Open("1.jpg")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	imgRef, err := vips.LoadImage(file)
	if err != nil {
		return nil, err
	}

	if err := vips.Webpsave(imgRef.Image(), image.Name()); err != nil {
		return nil, err
	}

	return image, nil
}
