package main

// Task represents a task entity
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	Deadline    string `json:"deadline"`
	Status      string `json:"status"`
}
