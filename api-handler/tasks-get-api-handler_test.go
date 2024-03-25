package api_handler

import (
	"net/http"
	"net/http/httptest"
	repository "tasks_api/repository/model"
	"testing"
)

func TestTasksGetAPIHandler(t *testing.T) {
	// Define test cases
	testCases := []struct {
		Name     string
		Method   string
		Path     string
		Expected int
	}{
		{
			Name:     "Valid ID",
			Method:   "GET",
			Path:     "/tasks/123",
			Expected: http.StatusOK,
		},
		{
			Name:     "Invalid ID",
			Method:   "GET",
			Path:     "/tasks/invalid",
			Expected: http.StatusBadRequest,
		},
	}

	// Create a new instance of TaskAPIHandler
	mockRepo := &MockTasksRepository{}
	handler := NewTaskAPIHandler(mockRepo)

	// Mock get response
	mockRepo.GetMock = func(ID int) (*repository.Task, error) {
		return &repository.Task{}, nil
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Create a new HTTP request
			req, err := http.NewRequest(tc.Method, tc.Path, nil)
			if err != nil {
				t.Fatal(err)
			}

			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()

			handler.TasksGetAPIHandler(rr, req)

			// Check the status code of the response
			if status := rr.Code; status != tc.Expected {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tc.Expected)
			}
		})
	}
}
