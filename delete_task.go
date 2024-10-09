package main

import (
	"fmt"
	"net/http"
)

func deleteTask(w http.ResponseWriter, r *http.Request, tasks *[]Task, id int) {
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
