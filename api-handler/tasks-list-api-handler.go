package api_handler

import (
	"log"
	"net/http"
)

func (h *TaskAPIHandler) TasksListAPIHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("Listing all tasks\n")
}
