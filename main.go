package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

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
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(message)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var task []Message

	if result := DB.Find(&task); result.Error != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var message Message

	taskID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := DB.First(&message, uint(taskID)).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	DB.Delete(&message, uint(taskID))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	taskID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var message Message
	var updates map[string]interface{}

	errP := json.NewDecoder(r.Body).Decode(&updates)
	if errP != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	updates["UpdatedAt"] = time.Now()

	if err := DB.First(&message, uint(taskID)).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	DB.Model(&message).Where("ID = ?", uint(taskID)).Updates(updates)
	DB.Save(&message)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func main() {

	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()

	router.HandleFunc("/api/messages", PostHandler).Methods("POST")
	router.HandleFunc("/api/messages", GetHandler).Methods("GET")
	router.HandleFunc("/api/messages/{id}", DeleteHandler).Methods("DELETE")
	router.HandleFunc("/api/messages/{id}", UpdateTaskHandler).Methods("PUT")

	http.ListenAndServe("localhost:8080", router)

}
