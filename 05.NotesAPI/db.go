package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	DB, err = gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting db: %v", err.Error())
	}

	DB.AutoMigrate(&Note{})
	fmt.Printf("Connected to DB successfully!\n")
}
