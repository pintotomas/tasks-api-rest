package api_handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (h *TaskAPIHandler) TasksDeleteAPIHandler(w http.ResponseWriter, r *http.Request) {
	// Split the URL path by '/'
	parts := strings.Split(r.URL.Path, "/")

	// Extract the ID from the URL path
	idStr := parts[2]

	// Convert the ID string to an integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	log.Printf("Deleting task with ID: %d\n", id)
}
