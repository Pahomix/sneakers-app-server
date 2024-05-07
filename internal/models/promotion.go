package models

import (
	"gorm.io/gorm"
	"time"
)

type Promotion struct {
	gorm.Model
	Name           string    `gorm:"not null" json:"name"`
	Code           string    `gorm:"unique;not null" json:"code"`
	ExpirationDate time.Time `gorm:"not null" json:"expiration_date"`
}
