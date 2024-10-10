package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getTaskID(w http.ResponseWriter, r *http.Request, tasks *[]Task, id int) {
	for _, task := range *tasks {
		if task.ID == id {
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Task not found")
}
