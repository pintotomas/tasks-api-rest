package api_handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

// TasksDeleteAPIHandler handles delete requests
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

	// Delete task
	err = h.tasksRepo.Delete(id)
	if err != nil {
		log.Printf("Error while deleting task: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
