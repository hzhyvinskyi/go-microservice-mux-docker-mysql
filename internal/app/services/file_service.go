package services

import (
	"context"
	"github.com/davidbyttow/govips/pkg/vips"
	"io/ioutil"
	"mime/multipart"
	"os"
)

const OutFilePath string = "tmp/out.webp"

// Upload creates temporary file. Opens it for reading and writing.
// Returns resulting *os.File.
func Upload(ctx context.Context, file multipart.File, header *multipart.FileHeader) (*os.File, error) {
	tempFile, err := ioutil.TempFile("tmp", "*-" + header.Filename)
	if err != nil {
		return nil, err
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if _, err := tempFile.Write(fileBytes); err != nil {
		return nil, err
	}

	return tempFile, nil
}

// Convert opens file by the given path.
// Loads an ImageRef from the given reader.
// Saves image in webp extension.
func Convert(ctx context.Context, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	imgRef, err := vips.LoadImage(file)
	if err != nil {
		return err
	}

	if err := vips.Webpsave(imgRef.Image(), OutFilePath); err != nil {
		return err
	}

	return nil
}

// Remove removes file from disk by given path.
func Remove(ctx context.Context, filePath string) error {
	if err := os.Remove(filePath); err != nil {
		return err
	}

	return nil
}
