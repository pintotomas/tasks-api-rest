package api_handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// TasksGetAPIHandler handles get requests
func (h *TaskAPIHandler) TasksGetAPIHandler(w http.ResponseWriter, r *http.Request) {
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

	// Get task
	task, err := h.tasksRepo.Get(id)
	if err != nil {
		log.Printf("Error while getting task: %s\n", err)
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	// Return task
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		log.Printf("Error while encoding task: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
