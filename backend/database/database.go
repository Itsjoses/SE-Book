package database

import (
	"fmt"
	"os"

	"github.com/Itsjoses/sebook-be/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// Read database connection string from environment variable
	dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Print(err)
	}

	DB = db

	DB.AutoMigrate(&models.User{}, &models.Genre{}, &models.Book{})
}
