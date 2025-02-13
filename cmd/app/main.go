package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"PetProject/internal/database"
	"PetProject/internal/handlers"
	"PetProject/internal/taskService"
	"PetProject/internal/web/tasks"
)

func main() {

	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(*repo)

	handler := handlers.NewHandler(service)

	// Инициализируем Echo
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start with err: %v", err)
	}

}
