package main

import (
	"PetProject/internal/database"
	"PetProject/internal/handlers"
	"PetProject/internal/taskService"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(*repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()

	router.HandleFunc("/api/messages", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/messages", handler.GetTaskHandler).Methods("GET")
	router.HandleFunc("/api/messages/{id}", handler.DeleteTaskHandler).Methods("DELETE")
	router.HandleFunc("/api/messages/{id}", handler.UpdateTaskHandler).Methods("PUT")

	http.ListenAndServe("localhost:8080", router)

}
