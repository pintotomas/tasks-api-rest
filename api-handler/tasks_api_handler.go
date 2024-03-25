package api_handler

import (
	"net/http"
	"regexp"
	"tasks_api/repository"
)

var (
	createTaskRequest = regexp.MustCompile(`^\/tasks[\/]*$`)
	updateTaskRequest = regexp.MustCompile(`^\/tasks[\/]*$`)
	getTaskRequest    = regexp.MustCompile(`^\/tasks\/(\d+)$`)
	listTasksRequest  = regexp.MustCompile(`^\/tasks[\/]*$`)
	deleteTaskRequest = regexp.MustCompile(`^\/tasks\/(\d+)$`)
)

type TaskAPIHandler struct {
	tasksRepo repository.TasksStorer
}

// NewTaskAPIHandler returns a new TaskAPIHandler
func NewTaskAPIHandler(tasksRepo repository.TasksStorer) *TaskAPIHandler {
	apiHandler := &TaskAPIHandler{}
	apiHandler.tasksRepo = tasksRepo
	return apiHandler
}

// ServeHTTP handles tasks http requests
func (h *TaskAPIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && listTasksRequest.MatchString(r.URL.Path):
		h.TasksListAPIHandler(w, r)
		return
	case r.Method == http.MethodGet && getTaskRequest.MatchString(r.URL.Path):
		h.TasksGetAPIHandler(w, r)
		return
	case r.Method == http.MethodPost && createTaskRequest.MatchString(r.URL.Path):
		h.TasksCreateAPIHandler(w, r)
		return
	case r.Method == http.MethodPut && updateTaskRequest.MatchString(r.URL.Path):
		h.TasksUpdateAPIHandler(w, r)
		return
	case r.Method == http.MethodDelete && deleteTaskRequest.MatchString(r.URL.Path):
		h.TasksDeleteAPIHandler(w, r)
		return
	default:
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
