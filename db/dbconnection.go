package db

import (
	"os"
	"path/filepath"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Connect - connect to the db
func Connect() error {
	// Load config for DB from os env vars. Defined in Docker Image
	// username, password, host, port
	log.Print("Connecting to DB...")
	db, err := sql.Open("mysql", "dataservice:password@/otw")
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Print("Connected!")
	migrateDB()
	return nil
}

func migrateDB() {
	// 1. Find all files with .sql endings. Must start with version_
	// 2. starting at one, load each file in turn and execute
	// 3. Sql files must not be destructive and must be idempotent
	log.Print("Working through SQL Files...")
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	
	matches, err := filepath.Glob(dir + "/db/*.sql")
	if err != nil {
		log.Fatal(err)
	}

	if len(matches) > 0 {
		log.Print("Found ", len(matches), " sql files to process!")
		for i := 0; i < len(matches); i++ {
			log.Print("Processing ", matches[i])
		}
	}
}
