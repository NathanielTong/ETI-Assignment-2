package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateTaskHandler handles the creation of a new task
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task Task

	// Parse the JSON-encoded task data from the request body
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Validate the task data
	if task.Title == "" {
		http.Error(w, "Task title is required", http.StatusBadRequest)
		return
	}

	// Insert the task into the database
	err = InsertTask(&task)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// UpdateTaskHandler handles the updating of a task
func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Extract task ID from URL path parameters
	id := extractTaskID(r)

	// Parse the JSON-encoded task data from the request body
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Validate the task data
	if task.Title == "" {
		http.Error(w, "Task title is required", http.StatusBadRequest)
		return
	}

	// Set the task ID
	task.ID = id

	// Update the task in the database
	err = UpdateTask(&task)
	if err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

// DeleteTaskHandler handles the deletion of a task
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Extract task ID from URL path parameters
	id := extractTaskID(r)

	// Delete the task from the database
	err := DeleteTask(id)
	if err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	// You can optionally return a success message or an empty response body
	// fmt.Fprintln(w, "Task deleted successfully")
}

// GetTaskHandler handles the retrieval of a task by ID
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Extract task ID from URL path parameters
	id := extractTaskID(r)

	// Retrieve the task from the database by ID
	task, err := GetTask(id)
	if err != nil {
		http.Error(w, "Failed to retrieve task", http.StatusInternalServerError)
		return
	}

	// Return the task data as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// ListTasksHandler handles the retrieval of all tasks with optional filtering
func ListTasksHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	queryParams := r.URL.Query()
	status := queryParams.Get("status")

	// Retrieve all tasks from the database
	tasks, err := ListTasks()
	if err != nil {
		http.Error(w, "Failed to retrieve tasks", http.StatusInternalServerError)
		return
	}

	// Filter tasks by status if provided
	if status != "" {
		filteredTasks := make([]*Task, 0)
		for _, task := range tasks {
			if task.Status == status {
				filteredTasks = append(filteredTasks, task)
			}
		}
		tasks = filteredTasks
	}

	// Return the list of tasks as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// extractTaskID extracts the task ID from URL path parameters
func extractTaskID(r *http.Request) int {
	vars := mux.Vars(r)
	idParam := vars["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// If the ID cannot be converted to an integer or is not provided, return a default value (e.g., 0)
		return 0
	}
	return id
}
