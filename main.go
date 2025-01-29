package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var task string

type requestBody struct {
	Task string `json:"task"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var req requestBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Inalid Request", http.StatusBadRequest)
		return
	}

	task = req.Task

	fmt.Fprintln(w, "Task update")
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", task)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/task", PostHandler).Methods("POST")
	router.HandleFunc("/", GetTaskHandler).Methods("GET")
	http.ListenAndServe("localhost:8080", router)

}
