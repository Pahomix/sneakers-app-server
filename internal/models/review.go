package models

import "gorm.io/gorm"

type Review struct {
	ID        uint            `gorm:"primaryKey" json:"id"`
	SneakerID uint            `gorm:"not null" json:"sneaker_id"`
	UserID    uint            `gorm:"not null" json:"user_id"`
	Rating    uint            `gorm:"not null" json:"rating"`
	Comment   string          `gorm:"not null" json:"comment"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"-"`
}
