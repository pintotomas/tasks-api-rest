package repository

import "time"

// Task represents the task model
type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`
	Status      int        `json:"status"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	Responsible string     `json:"responsible"`
	CreatedDate time.Time  `json:"created_date"`
	UpdatedDate time.Time  `json:"updated_date"`
}
