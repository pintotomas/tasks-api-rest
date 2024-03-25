package api_handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
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
	handler := TaskAPIHandler{}

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
