package app

import (
	"context"
	"github.com/davidbyttow/govips/pkg/vips"
	"log"
	"os"
)

func Convert(ctx context.Context) (string, error) {
	image := "out.webp"

	file, err := os.Open("1.jpg")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	imgRef, err := vips.LoadImage(file)
	if err != nil {
		log.Fatalln(err)
	}

	if err := vips.Webpsave(imgRef.Image(), image); err != nil {
		log.Fatalln(err)
	}

	return image, nil
}
