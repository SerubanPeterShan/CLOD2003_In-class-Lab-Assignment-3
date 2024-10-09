package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getAllTasks(w http.ResponseWriter, r *http.Request, tasks []Task) {
	err := json.NewDecoder(r.Body).Decode(&tasks)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if len(tasks) == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No tasks found")
		return
	}
	json.NewEncoder(w).Encode(tasks)

}
