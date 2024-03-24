package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Initialize db instance
	dbConn := InitializeDB()

	err := dbConn.Ping()
	if err != nil {
		log.Fatalf("Error checking db connection: %s\n", err)
		return
	}

	defer func(dbConn *sql.DB) {
		err = dbConn.Close()
		if err != nil {
			log.Fatalf("Error closing db connection: %s\n", err)
		}
	}(dbConn)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	http.ListenAndServe(":8080", nil)
}
