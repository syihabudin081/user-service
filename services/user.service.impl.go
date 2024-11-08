package services

import (
	"user-service/models"
	"user-service/repositories"
)

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *userService {
	return &userService{repo}
}

func (us *userService) CreateUser(user *models.User) error {
	return us.repo.CreateUser(user)
}

func (us *userService) GetAllUsers() ([]*models.User, error) {
	return us.repo.GetAllUsers()
}

func (us *userService) GetUserById(id string) (*models.User, error) {
	return us.repo.GetUserById(id)
}

func (us *userService) UpdateUser(user *models.User) error {
	return us.repo.UpdateUser(user)
}

func (us *userService) DeleteUser(id string) error {
	return us.repo.DeleteUser(id)
}
