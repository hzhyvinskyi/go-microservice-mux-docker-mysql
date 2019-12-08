package http

import (
	"github.com/hzhyvinskyi/webp-converter/internal/app/services"
	"github.com/hzhyvinskyi/webp-converter/pkg/errors"
	"net/http"
)

func ConvertHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		restError := errors.RestError{
			Message: "Max file size is 10MB",
			Status:	http.StatusBadRequest,
			Error:	"bad_request",
		}
		w.WriteHeader(restError.Status)
		w.Write([]byte(restError.Message))
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		restError := errors.RestError{
			Message: "Error while getting file",
			Status:	http.StatusBadRequest,
			Error:	"bad_request",
		}
		w.WriteHeader(restError.Status)
		w.Write([]byte(restError.Message))
		return
	}
	defer file.Close()

	image, err := services.Upload(file, header)
	if err != nil {
		restError := errors.RestError{
			Message: "Error while uploading file",
			Status:	http.StatusBadRequest,
			Error:	"bad_request",
		}
		w.WriteHeader(restError.Status)
		w.Write([]byte(restError.Message))
		return
	}

	converted, err := services.Convert(ctx, image)
	if err != nil {
		restError := errors.RestError{
			Message: "Error while converting",
			Status:	http.StatusBadRequest,
			Error:	"bad_request",
		}
		w.WriteHeader(restError.Status)
		w.Write([]byte(restError.Message))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "image/webp")
	http.ServeFile(w, r, converted.Name())
}
