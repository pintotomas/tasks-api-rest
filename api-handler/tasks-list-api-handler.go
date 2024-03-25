package api_handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

const (
	defaultPage     = 1
	defaultPageSize = 10
)

// TasksListAPIHandler handles list requests
func (h *TaskAPIHandler) TasksListAPIHandler(w http.ResponseWriter, r *http.Request) {

	// Parse query parameters
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = defaultPage
	}
	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil || size < 1 {
		size = defaultPageSize
	}

	// List tasks
	tasks, err := h.tasksRepo.List(page, size)
	if err != nil {
		log.Printf("Error while listing tasks: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return tasks
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		log.Printf("Error while encoding tasks: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
