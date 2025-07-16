package config

import (
	"fmt"
	"log"
	"os"

	"14.ContactListAPI/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	fmt.Println("Connecting db...")

	dbURL := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PORT"),
	)
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting db: %v", err.Error())
	}

	DB.AutoMigrate(&models.Contact{})

	fmt.Println("Connected db successfully!")
}
