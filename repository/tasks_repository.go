package repository

import (
	"database/sql"
	"fmt"
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
	defer closeStmt(stmt)

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
	// Prepare the SQL statement for retrieving a task by its ID
	stmt, err := r.db.Prepare(`
		SELECT ID, Title, Description, Status, DueDate, Responsible, CreatedDate, UpdatedDate
		FROM tasks
		WHERE ID = ?
	`)
	if err != nil {
		return nil, err
	}
	defer closeStmt(stmt)

	// Execute the SQL statement to fetch the task from the database
	row := stmt.QueryRow(ID)

	// Scan the database row and populate the task model
	task := &repository.Task{}
	err = row.Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.DueDate,
		&task.Responsible,
		&task.CreatedDate,
		&task.UpdatedDate,
	)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// Update updates an existing task in the database
func (r *TasksRepository) Update(task *model.UpdateTask) (*repository.Task, error) {
	// Prepare the SQL statement
	stmt, err := r.db.Prepare(`
		UPDATE tasks
		SET Title = ?, Description = ?, Status = ?, DueDate = ?, Responsible = ?, UpdatedDate = CURRENT_TIMESTAMP
		WHERE ID = ?
	`)
	if err != nil {
		return nil, err
	}
	defer closeStmt(stmt)

	// Execute the SQL statement for updating
	result, err := stmt.Exec(task.Title, task.Description, task.Status, task.DueDate, task.Responsible, task.ID)
	if err != nil {
		return nil, err
	}

	// Check the number of rows affected by the update operation
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	// If no rows were affected, return an error indicating that the task was not found
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	// Fetch the updated task from the database to ensure we have the latest data
	updatedTask, err := r.Get(task.ID)
	if err != nil {
		return nil, err // Return error if unable to fetch the updated task
	}

	// Return the updated task
	return updatedTask, nil
}

// Delete deletes a task from the database by its ID
func (r *TasksRepository) Delete(ID int) error {
	// Prepare the SQL statement
	stmt, err := r.db.Prepare("DELETE FROM tasks WHERE ID = ?")
	if err != nil {
		return err
	}
	defer closeStmt(stmt)

	// Execute the SQL statement for deleting
	_, err = stmt.Exec(ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *TasksRepository) List(page, size int) ([]*repository.Task, error) {
	// Calculate the offset based on the page and size parameters
	offset := (page - 1) * size

	// Prepare the SQL statement with LIMIT and OFFSET clauses
	stmt, err := r.db.Prepare(fmt.Sprintf("SELECT ID, Title, Description, Status, DueDate, Responsible, CreatedDate, UpdatedDate FROM tasks LIMIT %d OFFSET %d", size, offset))
	if err != nil {
		return nil, err
	}
	defer closeStmt(stmt)

	// Execute the SQL statement to query tasks with pagination
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	// Close the rows to release resources after scanning
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	// Initialize a slice to store the retrieved tasks
	var tasks []*repository.Task

	// Iterate through the rows of the result set and store each task in the result slice
	for rows.Next() {
		var task repository.Task

		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.DueDate, &task.Responsible, &task.CreatedDate, &task.UpdatedDate)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}

	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// closes the statement to free resources
func closeStmt(stmt *sql.Stmt) {
	err := stmt.Close()
	if err != nil {
		log.Printf("Error while closing Stmt: %s\n", err)
	}
}
