package main

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model

	Title   string `json:"title" gorm:"not null"`
	Content string `json:"content" gorm:"not null"`
}
