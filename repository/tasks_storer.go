package repository

import (
	"tasks_api/model"
	repository "tasks_api/repository/model"
)

// TasksStorer interface for managing tasks store
type TasksStorer interface {
	Create(task *model.Task) (*repository.Task, error)
	Get(ID int) (*repository.Task, error)
	Update(task *model.UpdateTask) (*repository.Task, error)
	Delete(ID int) error
	List(page, size int) ([]*repository.Task, error)
}
