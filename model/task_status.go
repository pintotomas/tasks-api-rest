package main

import "fmt"

// TaskStatus represents the status of a task
type TaskStatus int

// Enum-like constants for task status
const (
	Pending TaskStatus = iota
	InProgress
	Finished
)

// String representation of TaskStatus
func (s TaskStatus) String() string {
	switch s {
	case Pending:
		return "pending"
	case InProgress:
		return "in progress"
	case Finished:
		return "finished"
	default:
		return fmt.Sprintf("Unknown status: %d", s)
	}
}
