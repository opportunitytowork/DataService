package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Handling /welcome")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Welcome you!")
	})

	r.HandleFunc("/welcome/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Printf("Handling /welcome request: %v", vars["name"])
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Welcome %v!", vars["name"])
	})

	log.Print("Starting DataService...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
