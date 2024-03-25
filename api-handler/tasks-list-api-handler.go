package api_handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *TaskAPIHandler) TasksListAPIHandler(w http.ResponseWriter, r *http.Request) {

	// Get task
	tasks, err := h.tasksRepo.List()
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
