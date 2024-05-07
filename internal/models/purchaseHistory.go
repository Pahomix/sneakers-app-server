package models

import (
	"gorm.io/gorm"
	"time"
)

type PurchaseHistory struct {
	gorm.Model
	OrderID      uint      `gorm:"not null" json:"order_id"`
	UserID       uint      `gorm:"not null" json:"user_id"`
	SneakerID    uint      `gorm:"not null" json:"sneaker_id"`
	PurchaseDate time.Time `gorm:"not null" json:"purchase_date"`
	TotalAmount  float64   `gorm:"not null" json:"total_amount"`
}
