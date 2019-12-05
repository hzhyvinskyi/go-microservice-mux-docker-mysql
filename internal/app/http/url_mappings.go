package http

import (
	"net/http"
)

func mapUrls(router *http.ServeMux) {
	router.HandleFunc("/v1/convert", ConvertHandler)
}
