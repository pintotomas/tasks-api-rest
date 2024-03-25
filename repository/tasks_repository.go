package repository

import (
	"database/sql"
	"log"
	"tasks_api/model"
	repository "tasks_api/repository/model"
	"time"
)

type TasksRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TasksRepository {

	taskRepository := &TasksRepository{
		db: db,
	}
	return taskRepository
}

// Create inserts a new task into the database
func (r *TasksRepository) Create(task *model.Task) (*repository.Task, error) {

	// Prepare the SQL statement for inserting a new task
	stmt, err := r.db.Prepare(`
		INSERT INTO tasks (Title, Description, Status, DueDate, Responsible, CreatedDate, UpdatedDate)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		log.Printf("An error occured while creating task: %s\n", err)
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		err = stmt.Close()
		if err != nil {
			log.Printf("Error while closing Stmt: %s\n", err)
		}
	}(stmt)

	createDate, updateDate := time.Now(), time.Now()

	// Execute the SQL statement with the task data
	result, err := stmt.Exec(
		task.Title,
		task.Description,
		task.Status,
		task.DueDate,
		task.Responsible,
		createDate,
		updateDate,
	)
	if err != nil {
		return nil, err
	}

	// Retrieve the ID of the newly inserted task
	taskID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Construct the repository model for the created task
	createdTask := &repository.Task{
		ID:          int(taskID),
		Title:       task.Title,
		Description: task.Description,
		Status:      int(task.Status),
		DueDate:     task.DueDate,
		Responsible: task.Responsible,
		CreatedDate: createDate,
		UpdatedDate: updateDate,
	}

	return createdTask, nil
}

// Get retrieves a task from the database by its ID
func (r *TasksRepository) Get(ID int) (*repository.Task, error) {
	return nil, nil
}

// Update updates an existing task in the database
func (r *TasksRepository) Update(task *model.Task) (*repository.Task, error) {
	return nil, nil
}

// Delete deletes a task from the database by its ID
func (r *TasksRepository) Delete(ID int) error {
	return nil
}

// List retrieves a list of all tasks from the database
func (r *TasksRepository) List() ([]*repository.Task, error) {
	return nil, nil
}
