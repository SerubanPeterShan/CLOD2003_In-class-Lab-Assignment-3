package main

import (
	"encoding/json"
	"net/http"
)

func updateTask(w http.ResponseWriter, r *http.Request, tasks *[]Task, id int) {

	var updatedTask Task

	for i, task := range *tasks {
		if task.ID == id {
			if updatedTask.Title != "" {
				(*tasks)[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				(*tasks)[i].Description = updatedTask.Description
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode((*tasks)[i])
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)

}
