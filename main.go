package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	api_handler "tasks_api/api-handler"
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

	tasksAPIHandler := &api_handler.TaskAPIHandler{}

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})
	mux.Handle("/tasks", tasksAPIHandler)
	mux.Handle("/tasks/", tasksAPIHandler)
	fmt.Println("listening on 8000")
	http.ListenAndServe("localhost:8080", mux)

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprint(w, "Hello, World!")
	//})
	//
	//http.ListenAndServe(":8080", nil)
}
