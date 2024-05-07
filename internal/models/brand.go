package models

import "gorm.io/gorm"

type Brand struct {
	gorm.Model
	Name        string `gorm:"not null" json:"name"`
	Country     string `gorm:"not null" json:"country"`
	Description string `gorm:"not null" json:"description"`
}
