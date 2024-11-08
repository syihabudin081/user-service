package migrations

import (
	"user-service/config"
	"user-service/models"
)

func Migrate() {
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("Could not migrate the database")
	}
}
