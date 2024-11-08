package repositories

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	"user-service/models"
	"user-service/repositories"
)

// Setup in-memory database for testing
func setupTestDB() (*gorm.DB, error) {
	dsn := "host=localhost port=5432 user=postgres password=123 dbname=user_db sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.User{})
	return db, nil
}

func TestCreateUser(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := repositories.NewUserRepository(db)

	user := &models.User{
		Name:     "John Doe",
		Password: "hashedPassword",
		Email:    "johndoe@example.com",
	}

	// Call repository method
	err = repo.CreateUser(user)

	// Assert that the user was created successfully
	assert.NoError(t, err)

	// Fetch the user from the database
	var fetchedUser models.User
	db.First(&fetchedUser, "email = ?", user.Email)

	// Assert that the fetched user matches the expected user
	assert.Equal(t, user.Name, fetchedUser.Name)
	assert.Equal(t, user.Email, fetchedUser.Email)
}
