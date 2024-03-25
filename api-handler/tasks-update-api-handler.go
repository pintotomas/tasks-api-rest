package api_handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"tasks_api/model"
)

// TasksUpdateAPIHandler handles update requests
func (h *TaskAPIHandler) TasksUpdateAPIHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request body into a Update Task struct
	var updateTask *model.UpdateTask
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

	// Update task
	updatedTask, err := h.tasksRepo.Update(updateTask)
	if err != nil {
		log.Printf("Error while updating task: %s\n", err)
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	// Return updated task
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(updatedTask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error while encoding task: %s\n", err)
		return
	}
}
