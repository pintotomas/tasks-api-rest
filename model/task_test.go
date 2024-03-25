package model

import (
	"strings"
	"testing"
	"time"
)

func TestTaskValidation(t *testing.T) {

	// TestCase represents a test case for task validation
	type TestCase struct {
		Name  string
		Task  Task
		Valid bool
	}

	testCases := []TestCase{
		{
			Name: "Simple valid task",
			Task: Task{
				Title:       "Title",
				Description: "Description",
				Status:      InProgress, // Valid status
				DueDate:     ptr(time.Now()),
				Responsible: "Tomas",
			},
			Valid: true,
		},
		{
			Name: "Task should be valid on border scenarios",
			Task: Task{
				Title:       generateString(titleLimit),
				Description: generateString(descriptionLimit),
				Status:      InProgress, // Valid status
				DueDate:     ptr(time.Now()),
				Responsible: generateString(responsibleLimit),
			},
			Valid: true,
		},
		{
			Name: "Task should be invalid if the Status doesnt exist",
			Task: Task{
				Title:       "Title",
				Description: "Description",
				Status:      TaskStatus(5), // Invalid status
				Responsible: "Tomas",
			},
			Valid: false,
		},
		{
			Name: "Task should be invalid if the Title is empty",
			Task: Task{
				Title:       "", // Invalid title
				Status:      Finished,
				Responsible: "Tomas",
			},
			Valid: false,
		},
		{
			Name: "Task should be invalid if the Title has more characters than expected",
			Task: Task{
				Title:       generateString(titleLimit + 1),
				Status:      Finished,
				Responsible: "Tomas",
			},
			Valid: false,
		},
		{
			Name: "Task should be invalid if the Description has more characters than expected",
			Task: Task{
				Title:       "Title",
				Description: generateString(descriptionLimit + 1), // Invalid Description
				Status:      Finished,
				Responsible: "Tomas",
			},
			Valid: false,
		},
		{
			Name: "Task should be invalid if the Responsible has more characters than expected",
			Task: Task{
				Title:       "Title",
				Description: generateString(descriptionLimit + 1), // Invalid Description
				Status:      Finished,
				Responsible: generateString(responsibleLimit + 1),
			},
			Valid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			if result := tc.Task.Validate(); result != tc.Valid {
				t.Errorf("Validation failed, expected Validate to be %v, got %v", tc.Valid, result)
			}
		})
	}
}

func TestTaskValidationErrors(t *testing.T) {

	// TestCase represents a test case for task validation errors
	type TestCase struct {
		Name           string
		Task           Task
		ExpectedErrors []string
	}

	testCases := []TestCase{
		{
			Name: "Simple valid task",
			Task: Task{
				Title:       "Title",
				Description: "Description",
				Status:      InProgress, // Valid status
				DueDate:     ptr(time.Now()),
				Responsible: "Tomas",
			},
			ExpectedErrors: []string{},
		},
		{
			Name: "Task should be valid on border scenarios",
			Task: Task{
				Title:       generateString(titleLimit),
				Description: generateString(descriptionLimit),
				Status:      InProgress, // Valid status
				DueDate:     ptr(time.Now()),
				Responsible: generateString(responsibleLimit),
			},
			ExpectedErrors: []string{},
		},
		{
			Name: "Task should have error if the Status doesn't exist",
			Task: Task{
				Title:       "Title",
				Description: "Description",
				Status:      TaskStatus(5), // Invalid status
				Responsible: "Tomas",
			},
			ExpectedErrors: []string{ErrInvalidStatus},
		},
		{
			Name: "Task should have error if the Title is empty",
			Task: Task{
				Title:       "", // Invalid title
				Status:      Finished,
				Responsible: "Tomas",
			},
			ExpectedErrors: []string{ErrTitleRequired},
		},
		{
			Name: "Task should have error if the Title has more characters than expected",
			Task: Task{
				Title:       generateString(titleLimit + 1),
				Status:      Finished,
				Responsible: "Tomas",
			},
			ExpectedErrors: []string{ErrTitleMaxLength},
		},
		{
			Name: "Task should have error if the Description has more characters than expected",
			Task: Task{
				Title:       "Title",
				Description: generateString(descriptionLimit + 1), // Invalid Description
				Status:      Finished,
				Responsible: "Tomas",
			},
			ExpectedErrors: []string{ErrDescriptionMaxLength},
		},
		{
			Name: "Task should have error if the Responsible has more characters than expected",
			Task: Task{
				Title:       "Title",
				Description: generateString(descriptionLimit), // Invalid Description
				Status:      Finished,
				Responsible: generateString(responsibleLimit + 1),
			},
			ExpectedErrors: []string{ErrResponsibleMaxLength},
		},
		{
			Name: "Task should have error if the Responsible and Description have more characters than expected",
			Task: Task{
				Title:       "Title",
				Description: generateString(descriptionLimit + 1), // Invalid Description
				Status:      Finished,
				Responsible: generateString(responsibleLimit + 1),
			},
			ExpectedErrors: []string{ErrDescriptionMaxLength, ErrResponsibleMaxLength},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := tc.Task.ValidationErrors()
			if !equal(result.Errors, tc.ExpectedErrors) {
				t.Errorf("Expected errors: %v\nActual errors: %v", tc.ExpectedErrors, result.Errors)
			}
		})
	}
}

// generateString creates a string of n letters
func generateString(n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteByte('a')
	}
	return sb.String()
}

// equal checks if two slices of strings are equal
func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// Ptr nolint
func ptr[T any](x T) *T {
	return &x
}
