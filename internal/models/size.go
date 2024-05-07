package models

import "gorm.io/gorm"

type Size struct {
	gorm.Model
	SneakerID uint   `gorm:"not null" json:"sneaker_id"`
	Size      string `gorm:"not null" json:"size"`
	Quantity  uint   `gorm:"not null" json:"quantity"`
}
