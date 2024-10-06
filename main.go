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

// The main function starts the server on port 8091.
func main() {
	http.HandleFunc("/tasks", UserHandler)
	http.HandleFunc("/tasks/", UserHandler)

	http.HandleFunc("/updatetasks/", UserHandler_4_update)

	log.Println("Starting server on :8091...")
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
			getTaskID(w, r, &task, id)
		} else {
			getAllTasks(w, r, task)
		}
	case http.MethodPut:
		updateTask(w, r, &task, id)
	case http.MethodDelete:
		deleteTask(w, r, &task, id)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
	}
}

// creatiing new handler for edit as method cannot be determined from satsk ID

func UserHandler_4_update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/updatetasks/")
	id, _ := strconv.Atoi(idStr)

	switch r.Method {

	case http.MethodPost:

		updateTask(w, r, &task, id)
		fmt.Println("welcome to put bitch ")
	case http.MethodGet:
		if id > 0 {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			fmt.Println("get block")
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			fmt.Println("Hello, World!")
		}
	case http.MethodPut:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		fmt.Println(" put block ")

	case http.MethodDelete:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		fmt.Println("delete block")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		fmt.Println("default blcom ")
	}
}
