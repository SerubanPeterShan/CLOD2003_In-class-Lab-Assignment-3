package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getAllTasks(w http.ResponseWriter, tasks []Task) {
	if len(tasks) == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No tasks found")
		return
	}
	json.NewEncoder(w).Encode(tasks)

}
