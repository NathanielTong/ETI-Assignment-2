package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// InitDB initializes the database connection
func InitDB() error {
	connectionString := "username:password@tcp(localhost:3306)/task_management_db"

	conn, err := sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	db = conn
	return nil
}

// InsertTask inserts a task into the database
func InsertTask(task *Task) error {
	_, err := db.Exec(`
        INSERT INTO tasks (title, description, priority, deadline, status)
        VALUES (?, ?, ?, ?, ?)
    `, task.Title, task.Description, task.Priority, task.Deadline, task.Status)
	if err != nil {
		return err
	}
	return nil
}

// UpdateTask updates a task in the database
func UpdateTask(task *Task) error {
	_, err := db.Exec(`
        UPDATE tasks
        SET title = ?, description = ?, priority = ?, deadline = ?, status = ?
        WHERE id = ?
    `, task.Title, task.Description, task.Priority, task.Deadline, task.Status, task.ID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTask deletes a task from the database
func DeleteTask(taskID int) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", taskID)
	if err != nil {
		return err
	}
	return nil
}

// GetTask retrieves a task from the database by its ID
func GetTask(taskID int) (*Task, error) {
	var task Task
	err := db.QueryRow("SELECT * FROM tasks WHERE id = ?", taskID).Scan(
		&task.ID, &task.Title, &task.Description, &task.Priority, &task.Deadline, &task.Status)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// ListTasks retrieves all tasks from the database, optionally filtered by status
func ListTasks(status string) ([]*Task, error) {
	var query string
	var args []interface{}

	if status != "" {
		query = "SELECT * FROM tasks WHERE status = ?"
		args = append(args, status)
	} else {
		query = "SELECT * FROM tasks"
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.Deadline, &task.Status)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}
