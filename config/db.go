package config

import (
	"fmt"
	"os"
	"sirius/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDB() {
	var err error

	// Get environment variables
	var (
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		dbname   = os.Getenv("DB_NAME")
		timezone = os.Getenv("DB_TIMEZONE")
	)

	// Build connection string
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&timezone=%s", user, password, host, port, dbname, timezone)

	// Connect to database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Handle errors
	if err != nil {
		panic("Make sure you have set the correct database credentials in .env file")
	}

	// Migrate the schemas
	DB.AutoMigrate(&models.User{})
}
