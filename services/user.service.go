package services

import "user-service/models"

type UserService interface {
	CreateUser(user *models.User) error
	GetAllUsers() ([]*models.User, error)
	GetUserById(id string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
}
