package main

import "gorm.io/gorm"

type Todo struct {
	gorm.Model

	Title     string `json:"title" gorm:"not null"`
	Completed bool   `json:"completed" gorm:"not null;default:false"`
}
