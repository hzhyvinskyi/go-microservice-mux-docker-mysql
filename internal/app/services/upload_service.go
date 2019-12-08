package services

import (
	"io/ioutil"
	"mime/multipart"
	"os"
)

func Upload(file multipart.File, header *multipart.FileHeader) (*os.File, error) {
	tempFile, err := ioutil.TempFile("tmp", header.Filename)
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
