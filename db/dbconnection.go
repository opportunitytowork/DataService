package db

import (
	"log"
)

// Connect - connect to the db
func Connect() error {
	log.Print("Connecting to DB...")
	log.Print("Connected!")
	return nil
}
