package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize database connection
	err := InitDB()
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}
	defer db.Close()

	// Initialize router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/tasks", CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", UpdateTaskHandler).Methods("PUT")
	r.HandleFunc("/tasks/{id}", DeleteTaskHandler).Methods("DELETE")
	r.HandleFunc("/tasks", ListTasksHandler).Methods("GET")

	// Start the server
	port := ":8080" // or any port you prefer
	log.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(port, r))
}
