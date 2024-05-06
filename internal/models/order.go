package models

import "gorm.io/gorm"

type Order struct {
	ID          uint            `gorm:"primaryKey" json:"id"`
	UserID      uint            `gorm:"not null" json:"user_id"`
	Status      string          `gorm:"not null" json:"status"`
	TotalAmount float64         `gorm:"not null" json:"total_amount"`
	DeletedAt   *gorm.DeletedAt `gorm:"index" json:"-"`
}
