package controllers

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"net/http/httptest"
	"testing"
	"user-service/controllers"
	"user-service/models"
)

// Mocking the UserService
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserService) GetAllUsers() ([]*models.User, error) {
	args := m.Called()
	return args.Get(0).([]*models.User), args.Error(1)
}

func (m *MockUserService) GetUserById(id string) (*models.User, error) {
	args := m.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) UpdateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserService) DeleteUser(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateUser(t *testing.T) {
	// Initialize mock and controller
	mockService := new(MockUserService)
	controller := controllers.NewUserController(mockService)

	// Mock the CreateUser method with a callback to check the hashed password
	mockService.On("CreateUser", mock.MatchedBy(func(u *models.User) bool {
		// Check that the password is hashed correctly
		err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte("password"))
		return err == nil
	})).Return(nil)

	// Create a new Fiber app and context
	app := fiber.New()
	app.Post("/api/users", controller.CreateUser)

	// Prepare request body as JSON
	reqBody := []byte(`{"name":"John Doe","password":"password","email":"johndoe@example.com"}`)
	req := httptest.NewRequest("POST", "/api/users", bytes.NewReader(reqBody))

	req.Header.Set("Content-Type", "application/json")

	// Test the request
	resp, err := app.Test(req)
	assert.Nil(t, err, "Expected no error, but got: %v", err)

	// Assert the expected behavior
	assert.Equal(t, 201, resp.StatusCode)
	mockService.AssertExpectations(t)
}
func TestGetAllUsers(t *testing.T) {
	// Initialize mock and controller
	mockService := new(MockUserService)
	controller := controllers.NewUserController(mockService)

	// Test data
	users := []*models.User{
		{Name: "John Doe", Email: "johndoe@example.com"},
		{Name: "Jane Doe", Email: "janedoe@example.com"},
	}

	// Mock the GetAllUsers method
	mockService.On("GetAllUsers").Return(users, nil)

	// Create a new Fiber app and context
	app := fiber.New()
	app.Get("/api/users", controller.GetAllUsers)

	req := httptest.NewRequest("GET", "/api/users", nil)
	resp, _ := app.Test(req)

	// Assert the expected behavior
	assert.Equal(t, 200, resp.StatusCode)
	mockService.AssertExpectations(t)
}
