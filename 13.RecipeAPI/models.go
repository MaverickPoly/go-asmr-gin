package main

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model

	Name        string         `json:"name" gorm:"not null"`
	Ingredients pq.StringArray `json:"ingredients" gorm:"type:text[]"`
	Steps       pq.StringArray `json:"steps" gorm:"type:text[]"`
}
