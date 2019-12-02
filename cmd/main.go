package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"DataService/handlers"
	"DataService/db"
)

func main() {
	// Connect to the database. Must connect before starting the service
	if err := db.Connect(); err != nil {
		log.Fatal("Unable to connect to DB...")
	}
	
	// Setup the router
	r := mux.NewRouter()
	r.HandleFunc("/welcome", handlers.WelcomeHandler)
	r.HandleFunc("/welcome/{name}", handlers.WelcomeWithNameHandler)

	// Log starting of the service
	log.Print("Started DataService!")
	log.Fatal(http.ListenAndServe(":8000", r))
}
