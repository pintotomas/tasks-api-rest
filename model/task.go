package model

import (
	"time"
)

const (
	titleLimit       = 50
	descriptionLimit = 2000
	responsibleLimit = 50
)

const (
	ErrTitleRequired        = "Title is required"
	ErrTitleMaxLength       = "Title must be less than or equal to 50 characters"
	ErrDescriptionMaxLength = "Description must be less than or equal to 2000 characters"
	ErrInvalidStatus        = "Invalid status"
	ErrResponsibleRequired  = "Responsible is required"
	ErrResponsibleMaxLength = "Responsible must be less than or equal to 50 characters"
)

type Task struct {
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`
	Status      TaskStatus `json:"status"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	Responsible string     `json:"responsible"`
}

// Validate validates if a task is valid
func (t *Task) Validate() bool {
	if t.Title == "" || len(t.Title) > titleLimit {
		return false
	}
	if len(t.Description) > descriptionLimit {
		return false
	}
	if t.Status != Pending && t.Status != InProgress && t.Status != Finished {
		return false
	}
	if t.Responsible == "" || len(t.Responsible) > responsibleLimit {
		return false
	}
	return true
}

// ValidationErrors holds the errors of Task model
type ValidationErrors struct {
	Errors []string
}

// ValidationErrors returns validation errors
func (t *Task) ValidationErrors() *ValidationErrors {
	errs := ValidationErrors{}

	if t.Title == "" {
		errs.Errors = append(errs.Errors, ErrTitleRequired)
	}
	if len(t.Title) > titleLimit {
		errs.Errors = append(errs.Errors, ErrTitleMaxLength)
	}
	if len(t.Description) > descriptionLimit {
		errs.Errors = append(errs.Errors, ErrDescriptionMaxLength)
	}
	if t.Status != Pending && t.Status != InProgress && t.Status != Finished {
		errs.Errors = append(errs.Errors, ErrInvalidStatus)
	}
	if t.Responsible == "" {
		errs.Errors = append(errs.Errors, ErrResponsibleRequired)
	}
	if len(t.Responsible) > responsibleLimit {
		errs.Errors = append(errs.Errors, ErrResponsibleMaxLength)
	}

	return &errs
}
