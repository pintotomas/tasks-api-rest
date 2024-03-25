package model

const ErrIDRequired = "ID is required and must be greater than 0"

type UpdateTask struct {
	ID int `json:"id"`
	Task
}

// Validate validates if UpdateTask is valid
func (u *UpdateTask) Validate() bool {
	return u.ID > 0 && u.Task.Validate()
}

// ValidationErrors returns validation errors
func (u *UpdateTask) ValidationErrors() *ValidationErrors {
	validationErrors := ValidationErrors{}

	if u.ID <= 0 {
		validationErrors.Errors = []string{ErrIDRequired}
	}

	taskValidationErrors := u.Task.ValidationErrors()

	validationErrors.Errors = append(validationErrors.Errors, taskValidationErrors.Errors...)

	return &validationErrors
}
