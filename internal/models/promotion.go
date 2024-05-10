package models

import (
	"gorm.io/gorm"
)

type Promotion struct {
	gorm.Model
	Name           string `gorm:"not null" json:"name"`
	Code           string `gorm:"unique;not null" json:"code"`
	ExpirationDate string `gorm:"not null" json:"expiration_date"`
}
