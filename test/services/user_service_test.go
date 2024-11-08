package services

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"user-service/models"
	"user-service/services"
)

// Mock Repository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) GetAllUsers() ([]*models.User, error) {
	args := m.Called()
	return args.Get(0).([]*models.User), args.Error(1)
}

func (m *MockUserRepository) GetUserById(id string) (*models.User, error) {
	args := m.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) DeleteUser(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// Test Service
func TestCreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := services.NewUserService(mockRepo)

	user := &models.User{
		Name:     "John Doe",
		Password: "hashedPassword",
		Email:    "johndoe@example.com",
	}

	// Mock the repository's CreateUser method to return nil (no error)
	mockRepo.On("CreateUser", user).Return(nil)

	// Call the service method
	err := userService.CreateUser(user)

	// Assert that the service method behaves correctly
	assert.NoError(t, err)

	// Assert that the repository method was called with the correct arguments
	mockRepo.AssertExpectations(t)
}

func TestCreateUserWithError(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := services.NewUserService(mockRepo)

	user := &models.User{
		Name:     "John Doe",
		Password: "hashedPassword",
		Email:    "johndoe@example.com",
	}

	// Mock the repository's CreateUser method to return an error
	mockRepo.On("CreateUser", user).Return(errors.New("error creating user"))

	// Call the service method
	err := userService.CreateUser(user)

	// Assert that the service method returns an error
	assert.Error(t, err)

	// Assert that the repository method was called
	mockRepo.AssertExpectations(t)
}
