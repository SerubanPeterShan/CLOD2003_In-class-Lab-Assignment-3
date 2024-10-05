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

func main() {
	http.HandleFunc("/tasks", UserHandler)
	http.HandleFunc("/tasks/", UserHandler)

	log.Println("Starting server on :8091...")
	if err := http.ListenAndServe(":8091", nil); err != nil {
		log.Fatal(err)
	}
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, _ := strconv.Atoi(idStr)

	switch r.Method {
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
