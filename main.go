package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	result := DB.Create(&message)
	if result.Error != nil {
		http.Error(w, "Failed to create message", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Task added to database")
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var task []Message

	if result := DB.Find(&task); result.Error != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(task)
}

func main() {

	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()

	router.HandleFunc("/api/messages", PostHandler).Methods("POST")
	router.HandleFunc("/api/messages", GetHandler).Methods("GET")

	http.ListenAndServe("localhost:8080", router)

}
