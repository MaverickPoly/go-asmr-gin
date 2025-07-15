package config

import (
	"fmt"
	"log"
	"os"

	"08.BookManagement/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	fmt.Println("Connecting db...")

	var err error
	dbUrl := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PORT"),
	)
	DB, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err.Error())
	}

	DB.AutoMigrate(&models.Book{})

	fmt.Println("Connected db!")
}
