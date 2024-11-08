package test

import (
	"bytes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http/httptest"
	"testing"
	"user-service/controllers"
	"user-service/models"
	"user-service/repositories"
	"user-service/services"
)

// setupTestDB initializes the database connection and ensures the necessary migrations are applied.
func setupTestDB() (*gorm.DB, error) {
	dsn := "host=localhost port=5432 user=postgres password=123 dbname=user_db sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	// Ensure that AutoMigrate doesn't return errors
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, fmt.Errorf("failed to run AutoMigrate: %v", err)
	}

	return db, nil
}

func TestCreateUserIntegration(t *testing.T) {
	// Initialize the real database connection (e.g., Postgres)
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Error setting up test database: %v", err)
	}

	// Initialize the repository, service, and controller
	repo := repositories.NewUserRepository(db)
	assert.NotNil(t, repo, "User repository should not be nil")

	userService := services.NewUserService(repo)
	assert.NotNil(t, userService, "User service should not be nil")

	controller := controllers.NewUserController(userService)
	assert.NotNil(t, controller, "User controller should not be nil")

	// Create a new Fiber app and context
	app := fiber.New()
	app.Post("/users", controller.CreateUser)

	// Prepare the test request with the Content-Type header
	req := httptest.NewRequest("POST", "/users", bytes.NewReader([]byte(`{"name":"test","password":"password1","email":"test@example.com"}`)))
	req.Header.Set("Content-Type", "application/json")

	// Send the request to the Fiber app
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error occurred while testing the request: %v", err)
	}

	// Assert that the user was created correctly (Status Code 201 Created)
	assert.Equal(t, 201, resp.StatusCode)

	// Check the database to see if the user exists
	var fetchedUser models.User
	result := db.First(&fetchedUser, "email = ?", "test@example.com")
	if result.Error != nil {
		t.Fatalf("Failed to fetch user from database: %v", result.Error)
	}

	// Assert the fetched user matches the input data
	assert.Equal(t, "test", fetchedUser.Name)
	assert.Equal(t, "test@example.com", fetchedUser.Email)
}
