package http

import (
	"net/http"
)

func mapUrls(router *http.ServeMux) {
	router.HandleFunc("/convert", ConvertHandler())
}
