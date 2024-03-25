package api_handler

import (
	"encoding/json"
	"log"
	"net/http"
	"tasks_api/model"
)

// TasksCreateAPIHandler handles creates requests
func (h *TaskAPIHandler) TasksCreateAPIHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request body into a Task struct
	var task *model.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		// If there's an error decoding the JSON body, return a 400 Bad Request response
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the task object
	if valid := task.Validate(); !valid {
		// Return error response
		errResponse := &ErrorResponse{Errors: task.ValidationErrors().Errors}
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(errResponse)
		if err != nil {
			return
		}
		return
	}

	// Save task
	createdTask, err := h.tasksRepo.Create(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error while creating task: %s\n", err)
		return
	}

	// Return created task
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdTask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error while encoding task: %s\n", err)
		return
	}
}
