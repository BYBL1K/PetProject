package main

import (
	"encoding/json"
	"net/http"
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
	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	if err := DB.First(&message, message.ID).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	DB.Delete(&message, message.ID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var message Message

	var updates map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	updates["UpdatedAt"] = time.Now()

	if err := DB.First(&message, updates["ID"]).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	DB.Model(&message).Where("id = ?", message.ID).Updates(updates)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func main() {

	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()

	router.HandleFunc("/api/messages", PostHandler).Methods("POST")
	router.HandleFunc("/api/messages", GetHandler).Methods("GET")
	router.HandleFunc("/api/messages", DeleteHandler).Methods("DELETE")
	router.HandleFunc("/api/messages", UpdateTaskHandler).Methods("PUT")

	http.ListenAndServe("localhost:8080", router)

}
