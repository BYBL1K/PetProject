package userService

import (
	"PetProject/internal/taskService"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Tasks    []taskService.Task
}
