package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"DataService/handlers"
)

func main() {
	// Setup the router
	r := mux.NewRouter()
	r.HandleFunc("/welcome", handlers.WelcomeHandler)
	r.HandleFunc("/welcome/{name}", handlers.WelcomeWithNameHandler)

	// Log starting of the service
	log.Print("Starting DataService...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
