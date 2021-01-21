package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/jtobin321/go-jwt-api/api/models"

	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres database driver
)

// InitializeDB initializes database connection for server
func InitializeDB() (*gorm.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dbURL := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		dbHost,
		dbPort,
		dbUser,
		dbName,
		dbPass,
	)

	DB, err := gorm.Open("postgres", dbURL)
	if err != nil {
		fmt.Printf("Cannot connect to database: %s", dbName)
		log.Fatal("This is the error:", err)
		return nil, err
	}

	fmt.Printf("Successfully connected to database: %s", dbName)

	// Run database migrations
	DB.Debug().AutoMigrate(&models.User{}, &models.Post{})

	return DB, nil
}
