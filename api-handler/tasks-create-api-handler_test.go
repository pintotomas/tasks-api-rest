package api_handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"tasks_api/model"
	repository "tasks_api/repository/model"
	"testing"
)

func TestTasksCreateAPIHandler(t *testing.T) {

	// Define a struct to represent a test case scenario
	type testCase struct {
		Name          string // Name of the test case
		Payload       []byte // JSON payload
		ExpectedCode  int    // Expected HTTP status code
		ExpectedError string // Expected error message
	}

	// Define test cases
	testCases := []testCase{
		{
			Name:          "Valid request",
			Payload:       []byte(`{"title":"Task title","description":"Task description","status":0,"responsible":"Tomas Pinto", "due_date":"2022-03-31T12:00:00Z"}`),
			ExpectedCode:  http.StatusCreated,
			ExpectedError: "", // No error expected
		},
		{
			Name:          "Empty title",
			Payload:       []byte(`{"title":"","description":"Task description","status":0,"responsible":"John Doe"}`),
			ExpectedCode:  http.StatusBadRequest,
			ExpectedError: "Title is required",
		},
		{
			Name:          "Invalid status",
			Payload:       []byte(`{"title":"Task title","description":"Task description","status":-1,"responsible":"John Doe"}`),
			ExpectedCode:  http.StatusBadRequest,
			ExpectedError: "Invalid status",
		},
		{
			Name:          "Empty responsible",
			Payload:       []byte(`{"title":"Task title","description":"Task description","status":1,"responsible":""}`),
			ExpectedCode:  http.StatusBadRequest,
			ExpectedError: "Responsible is required",
		},
		{
			Name:          "Very long title",
			Payload:       []byte(`{"title":"LongTitleLongTitleLongTitleLongTitleLongTitleLongTitle","description":"Task description","status":1,"responsible":"Tomas Pinto"}`),
			ExpectedCode:  http.StatusBadRequest,
			ExpectedError: "Title must be less than or equal to 50 characters",
		},
		{
			Name:          "Very long title, missing responsible",
			Payload:       []byte(`{"title":"LongTitleLongTitleLongTitleLongTitleLongTitleLongTitle","description":"Task description","status":1,"responsible":""}`),
			ExpectedCode:  http.StatusBadRequest,
			ExpectedError: `"Title must be less than or equal to 50 characters","Responsible is required"`,
		},
	}

	// Create a new instance of TaskAPIHandler
	mockRepo := &MockTasksRepository{}
	handler := NewTaskAPIHandler(mockRepo)

	// Mock create response
	mockRepo.CreateMock = func(task *model.Task) (*repository.Task, error) {
		return &repository.Task{}, nil
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Create a new HTTP request with the JSON payload
			req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(tc.Payload))
			if err != nil {
				t.Fatal(err)
			}

			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()

			handler.TasksCreateAPIHandler(rr, req)

			// Check the status code of the response
			if status := rr.Code; status != tc.ExpectedCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.ExpectedCode)
			}

			// Check the response body to ensure it contains the appropriate error message
			if tc.ExpectedError != "" {
				if body := rr.Body.String(); !strings.Contains(body, tc.ExpectedError) {
					t.Errorf("handler returned unexpected body: got %v want %v",
						body, tc.ExpectedError)
				}
			}
		})
	}
}

type MockTasksRepository struct {
	CreateMock func(task *model.Task) (*repository.Task, error)
	GetMock    func(ID int) (*repository.Task, error)
	UpdateMock func(task *model.Task) (*repository.Task, error)
	DeleteMock func(ID int) error
	ListMock   func() ([]*repository.Task, error)
}

func (m MockTasksRepository) Create(task *model.Task) (*repository.Task, error) {
	if m.CreateMock == nil {
		return nil, errors.New("mock for Create does not exist")
	}
	return m.CreateMock(task)
}

func (m MockTasksRepository) Get(ID int) (*repository.Task, error) {
	if m.GetMock == nil {
		return nil, errors.New("mock for Get does not exist")
	}
	return m.Get(ID)
}

func (m MockTasksRepository) Update(task *model.Task) (*repository.Task, error) {
	if m.UpdateMock == nil {
		return nil, errors.New("mock for Update does not exist")
	}
	return m.Update(task)
}

func (m MockTasksRepository) Delete(ID int) error {
	if m.DeleteMock == nil {
		return errors.New("mock for Delete does not exist")
	}
	return m.DeleteMock(ID)
}

func (m MockTasksRepository) List() ([]*repository.Task, error) {
	if m.ListMock == nil {
		return nil, errors.New("mock for List does not exist")
	}
	return m.ListMock()
}
