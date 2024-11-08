package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func DatabaseInit() {

	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	dsn := os.Getenv("DB_URL")
	var retries = 5
	for retries > 0 {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println("Database connection failed. Retrying...")
			retries--
		} else {
			fmt.Println("Database connected")
			return
		}
		log.Fatalf("Failed to connect to database: %v", err)
	}
}
