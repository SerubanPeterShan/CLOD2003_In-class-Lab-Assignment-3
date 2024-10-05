package main

import (
	"encoding/json"
	"net/http"
)

func createTask(w http.ResponseWriter, r *http.Request, tasks *[]Task, nextID *int) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Assign the next available ID to the task
	task.ID = *nextID
	// Increment the nextID for the next task to global variable
	*nextID++
	// Append the task to the tasks slice to golbal variable(to simulate a database)
	*tasks = append(*tasks, task)
	json.NewEncoder(w).Encode(task)

}
