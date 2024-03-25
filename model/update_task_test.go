package model

import (
	"testing"
	"time"
)

func TestUpdateTaskValidation(t *testing.T) {
	// TestCase represents a test case for UpdateTask validation
	type TestCase struct {
		Name       string
		UpdateTask UpdateTask
		Valid      bool
	}

	testCases := []TestCase{
		{
			Name: "Valid UpdateTask",
			UpdateTask: UpdateTask{
				ID: 123,
				Task: Task{
					Title:       "Title",
					Description: "Description",
					Status:      InProgress,
					DueDate:     ptr(time.Now()),
					Responsible: "Tomas",
				},
			},
			Valid: true,
		},
		{
			Name: "UpdateTask missing ID",
			UpdateTask: UpdateTask{
				ID: 0, // Invalid ID
				Task: Task{
					Title:       "Title",
					Description: "Description",
					Status:      InProgress,
					DueDate:     ptr(time.Now()),
					Responsible: "Tomas",
				},
			},
			Valid: false,
		},
		{
			Name: "UpdateTask with invalid Task",
			UpdateTask: UpdateTask{
				ID: 123,
				Task: Task{
					Title:  "", // Invalid title
					Status: Finished,
				},
			},
			Valid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := tc.UpdateTask.Validate()
			if result != tc.Valid {
				t.Errorf("Expected test validation to be: %v", tc.Valid)
			}
		})
	}
}
