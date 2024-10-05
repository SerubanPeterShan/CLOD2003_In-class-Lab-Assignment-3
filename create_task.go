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
	task.ID = *nextID
	*nextID++
	*tasks = append(*tasks, task)
	json.NewEncoder(w).Encode(task)

}
