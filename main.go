package main

import (
	"log"
	"net/http"
	"go-microservice-mux-docker-mysql/db"
	"github.com/gorilla/mux"
)

func setupRouter(router *mux.Router) {
	router.Methods("POST").Path("/endpoint").HandlerFunc(postHandler)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	database, err := db.CreateDatabase()
	if err != nil {
		log.Fatal("Database connection failed!")
	}

	_, err = database.Exec("INSERT INTO `test` (name) VALUES ('testName')")
	if err != nil {
		log.Fatal("Database QUERY failed!")
	}

	log.Println("You called resource!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	setupRouter(router)
	log.Fatal(http.ListenAndServe(":9092", router))
}
