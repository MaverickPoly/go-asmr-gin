package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	Title         string `json:"title" gorm:"not null"`
	Author        string `json:"author" gorm:"not null"`
	PublishedYear uint   `json:"published_year" gorm:"not null"`
}
