package http

import (
	"github.com/hzhyvinskyi/webp-converter/internal/app"
	"log"
	"net/http"
)

func ConvertHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	result, err := app.Convert(ctx)
	if err != nil {
		log.Fatalln(err)

	}

	w.WriteHeader(http.StatusCreated)
	http.ServeFile(w, r, result)
}
