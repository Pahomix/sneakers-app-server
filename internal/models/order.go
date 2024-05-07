package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID      uint    `gorm:"not null" json:"user_id"`
	Status      string  `gorm:"not null" json:"status"`
	TotalAmount float64 `gorm:"not null" json:"total_amount"`
}
