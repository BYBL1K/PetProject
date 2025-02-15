package userService

import (
	"PetProject/internal/taskService"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user User) (User, error)

	GetAllUsers() ([]User, error)

	UpdateUserByID(id uint, user User) (User, error)

	DeleteUserByID(id uint) error

	GetTasksForUser(userID uint) ([]taskService.Task, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) UpdateUserByID(id uint, updatedUser User) (User, error) {
	var user User

	err := r.db.First(&user, id).Error
	if err != nil {
		return user, err
	}

	r.db.Model(&user).Where("ID = ?", id).Updates(updatedUser)

	err = r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil

}

func (r *userRepository) DeleteUserByID(id uint) error {
	var user User

	err := r.db.First(&user, id).Error
	if err != nil {
		return err
	}

	err = r.db.Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetTasksForUser(userID uint) ([]taskService.Task, error) {
	var tasks []taskService.Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}
