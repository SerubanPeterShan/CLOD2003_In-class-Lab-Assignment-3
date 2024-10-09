package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func deleteTask(w http.ResponseWriter, r *http.Request, tasks *[]Task, id int) {
	err := json.NewDecoder(r.Body).Decode(&tasks)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	for i, task := range *tasks {
		if task.ID == id {
			*tasks = append((*tasks)[:i], (*tasks)[i+1:]...)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Task with ID %d deleted", id)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Task with ID %d not found", id)
}
