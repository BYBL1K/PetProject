package userService

import "PetProject/internal/taskService"

type UserService struct {
	repo userRepository
}

func NewService(repo userRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user User) (User, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) UpdateUserByID(id uint, updatedUser User) (User, error) {
	return s.repo.UpdateUserByID(id, updatedUser)
}

func (s *UserService) DeleteUserByID(id uint) error {
	return s.repo.DeleteUserByID(id)
}

func (s *UserService) GetTasksForUser(userID uint) ([]taskService.Task, error) {
	return s.repo.GetTasksForUser(userID)
}
