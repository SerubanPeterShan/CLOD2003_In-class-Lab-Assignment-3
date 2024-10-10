package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func deleteTask(w http.ResponseWriter, tasks *[]Task, id int) {
	for i, task := range *tasks {
		if task.ID == id {
			*tasks = append((*tasks)[:i], (*tasks)[i+1:]...)
			w.WriteHeader(http.StatusOK)
			response := Response{Message: fmt.Sprintf("Task with ID %d deleted", id)}
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	response := Response{Message: fmt.Sprintf("Task with ID %d not found", id)}
	json.NewEncoder(w).Encode(response)
}
