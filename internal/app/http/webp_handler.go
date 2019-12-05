package http

import (
	"encoding/json"
	"github.com/hzhyvinskyi/webp-converter/internal/app"
	"log"
	"net/http"
)

func ConvertHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request catched by ConvertHandler")
		ctx := r.Context()

		result, err := app.Convert(ctx)
		if err != nil {
			log.Fatalln(err)
		}

		if err := json.NewEncoder(w).Encode(result); err != nil {
			log.Fatalln(err)
		}
	}
}
