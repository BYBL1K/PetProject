package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"PetProject/internal/database"
	"PetProject/internal/handlers"
	"PetProject/internal/taskService"
	"PetProject/internal/userService"
	"PetProject/internal/web/tasks"
	"PetProject/internal/web/users"
)

func main() {

	database.InitDB()
	errTask := database.DB.AutoMigrate(&taskService.Task{})
	if errTask != nil {
		log.Fatalf("Failed to AutoMigrate Database with err: %v", errTask)
	}

	errUser := database.DB.AutoMigrate(&userService.User{})
	if errUser != nil {
		log.Fatalf("Failed to AutoMigrate Database with err: %v", errUser)
	}

	tasksRepo := taskService.NewTaskRepository(database.DB)
	usersRepo := userService.NewUserRepository(database.DB)

	tasksService := taskService.NewService(*tasksRepo)
	usersService := userService.NewService(*usersRepo)

	tasksHandler := handlers.NewTaskHandler(tasksService)
	usersHandler := handlers.NewUserHandler(usersService)

	// Инициализируем Echo
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictTasksHandler := tasks.NewStrictHandler(tasksHandler, nil)
	tasks.RegisterHandlers(e, strictTasksHandler)

	strictUsersHandler := users.NewStrictHandler(usersHandler, nil)
	users.RegisterHandlers(e, strictUsersHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start with err: %v", err)
	}

}
