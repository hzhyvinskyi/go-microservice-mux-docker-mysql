package http

import (
	"github.com/hzhyvinskyi/webp-converter/internal/app/services"
	"github.com/hzhyvinskyi/webp-converter/pkg/errors"
	"net/http"
)

const MaxMemory = 10 << 20

func ConvertHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := r.ParseMultipartForm(MaxMemory); err != nil {
		errors.Response(w, "Max file size is 10MB", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		errors.Response(w, "File retrieving error", http.StatusBadRequest)
		return
	}
	defer file.Close()

	tempFile, err := services.Upload(ctx, file, header)
	if err != nil {
		errors.Response(w, "File upload error", http.StatusBadRequest)
		return
	}

	if err = services.Convert(ctx, tempFile.Name()); err != nil {
		errors.Response(w, "File convert error", http.StatusBadRequest)
		return
	}

	if err = services.Remove(ctx, tempFile.Name()); err != nil {
		errors.Response(w, "File remove error", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "image/webp")
	http.ServeFile(w, r, services.OutFilePath)
}
