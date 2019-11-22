package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// WelcomeHandler - Handler
func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling /welcome")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome you!")
}

// WelcomeWithNameHandler - Handler
func WelcomeWithNameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("Handling /welcome request: %v", vars["name"])
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome %v!", vars["name"])
}
