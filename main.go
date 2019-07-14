package main

import (
	"log"
	"net/http"

	"go-microservice-mux-docker-mysql/app"
	"go-microservice-mux-docker-mysql/db"
	"github.com/gorilla/mux"
)

func main() {
	database, err := db.CreateDatabase()
	if err != nil {
		log.Fatal("Database connection failed: %s", err.Error())
	}

	app := &app.App {
		mux.NewRouter().StrictSlash(true),
		database,
	}

	app.SetupRouter()

	log.Fatal(http.ListenAndServe(":9092", app.Router))
}
