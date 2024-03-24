package main

import "testing"

func TestTaskStatusString(t *testing.T) {
	testCases := []struct {
		Status   TaskStatus
		Expected string
	}{
		{Status: Pending, Expected: "pending"},
		{Status: InProgress, Expected: "in progress"},
		{Status: Finished, Expected: "finished"},
		{Status: TaskStatus(100), Expected: "Unknown status: 100"},
	}

	for _, tc := range testCases {
		t.Run(tc.Expected, func(t *testing.T) {
			if result := tc.Status.String(); result != tc.Expected {
				t.Errorf("Expected status string %s, but got %s", tc.Expected, result)
			}
		})
	}
}
