package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"` // "pending" or "completed"
}

// In-memory storage (simulating a database)
var task = []Task{}
var nextID = 1

// The main function starts the server on port 8092.
func main() {
	http.HandleFunc("/tasks", UserHandler)
	http.HandleFunc("/tasks/", UserHandler)

	log.Println("Starting server on :8092...")
	if err := http.ListenAndServe(":8092", nil); err != nil {
		log.Fatal(err)
	}
}

// This function is the handler for the /tasks and /tasks/ endpoints.
// It parses the URL to determine if the request is for a specific task or all tasks.
// It then delegates the request to the appropriate handler function.
func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, _ := strconv.Atoi(idStr)

	switch r.Method {
	//&task is a pointer to the task slice and &nextID is a pointer to the nextID variable
	case http.MethodPost:
		createTask(w, r, &task, &nextID)
	case http.MethodGet:
		if id > 0 {
			getTaskID(w, &task, id)
		} else {
			getAllTasks(w, task)
		}
	case http.MethodPut:
		updateTask(w, r, &task, id)
	case http.MethodDelete:
		deleteTask(w, &task, id)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
	}
}
