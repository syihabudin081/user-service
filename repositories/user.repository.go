package repositories

import (
	"gorm.io/gorm"
	"user-service/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetAllUsers() ([]*models.User, error)
	GetUserById(id string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) CreateUser(user *models.User) error {
	if err := ur.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	if err := ur.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) GetUserById(id string) (*models.User, error) {
	user := new(models.User)
	if err := ur.DB.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) UpdateUser(user *models.User) error {
	if err := ur.DB.Model(&models.User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) DeleteUser(id string) error {
	if err := ur.DB.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}
