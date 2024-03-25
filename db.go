package main

import (
	"database/sql"
	// sqlite driver
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func InitializeDB() *sql.DB {
	// Open a database connection
	// TODO improvement: move to env variables
	db, err := sql.Open("sqlite3", "task.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create tasks table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
    id INTEGER PRIMARY KEY,
    Title VARCHAR(50) NOT NULL,
    Description VARCHAR(2000),
    Status INTEGER CHECK(Status >= 0 AND Status <= 2) NOT NULL,
    DueDate DATE,
    Responsible VARCHAR(50) NOT NULL,
    CreatedDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)`)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
