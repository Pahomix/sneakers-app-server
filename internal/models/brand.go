package models

import "gorm.io/gorm"

type Brand struct {
	ID          uint            `gorm:"primaryKey" json:"id"`
	Name        string          `gorm:"not null" json:"name"`
	Country     string          `gorm:"not null" json:"country"`
	Description string          `gorm:"not null" json:"description"`
	DeletedAt   *gorm.DeletedAt `gorm:"index" json:"-"`
}
