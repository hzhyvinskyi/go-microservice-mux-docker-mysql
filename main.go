package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func setupRouter(router *mux.Router) {
	router.Methods("POST").Path("/endpoint").HandlerFunc(postHandler)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("You called resource!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	setupRouter(router)
	log.Fatal(http.ListenAndServe(":9092", router))
}
