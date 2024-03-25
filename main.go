package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	api_handler "tasks_api/api-handler"
	repository2 "tasks_api/repository"
)

func main() {

	// Initialize db instance
	dbConn := InitializeDB()

	err := dbConn.Ping()
	if err != nil {
		log.Fatalf("Error checking db connection: %s\n", err)
		return
	}

	// Close db connection when we stop the API
	defer func(dbConn *sql.DB) {
		err = dbConn.Close()
		if err != nil {
			log.Fatalf("Error closing db connection: %s\n", err)
		}
	}(dbConn)

	// Initialize
	repository := repository2.NewTaskRepository(dbConn)

	tasksAPIHandler := api_handler.NewTaskAPIHandler(repository)

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})
	mux.Handle("/tasks", tasksAPIHandler)
	mux.Handle("/tasks/", tasksAPIHandler)
	fmt.Println("listening on 8000")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Error while initializing server: %s\n", err)
		return
	}

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprint(w, "Hello, World!")
	//})
	//
	//http.ListenAndServe(":8080", nil)
}
