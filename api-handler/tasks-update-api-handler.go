package api_handler

import (
	"encoding/json"
	"net/http"
	"tasks_api/model"
)

func (h *TaskAPIHandler) TasksUpdateAPIHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request body into a Update Task struct
	var updateTask model.UpdateTask
	if err := json.NewDecoder(r.Body).Decode(&updateTask); err != nil {
		// If there's an error decoding the JSON body, return a 400 Bad Request response
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the task object
	if valid := updateTask.Validate(); !valid {
		// Return error response
		errResponse := &ErrorResponse{Errors: updateTask.ValidationErrors().Errors}
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(errResponse)
		if err != nil {
			return
		}
		return
	}

	// TODO connect with repository to save the task

	w.WriteHeader(http.StatusOK)
}
