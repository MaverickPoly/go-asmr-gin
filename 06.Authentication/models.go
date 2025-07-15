package main

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string `json:"username" gorm:"not null;unique"`
	Email    string `json:"email" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
	Location string `json:"location" gorm:"not null"`
}
