package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	fmt.Println("Connecting db...")
	var err error

	db, err = gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect sqlite - %v", err.Error())
	}
	db.AutoMigrate(&User{})

	fmt.Println("Connected to sqlite database successfully!")
}
