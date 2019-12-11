package http

import (
	"github.com/gorilla/mux"
)

func mapUrls(router *mux.Router) {
	v1 := router.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/convert", ConvertHandler).Methods("POST")
}
