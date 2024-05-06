package models

import "gorm.io/gorm"

type User struct {
	ID        uint            `gorm:"primaryKey" json:"id"`
	Username  string          `gorm:"not null" json:"username"`
	Email     string          `gorm:"unique;not null" json:"email"`
	Password  string          `gorm:"not null" json:"-"`
	Role      string          `gorm:"not null" json:"role"`
	UserPhoto *string         `gorm:"-" json:"user_photo,omitempty"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"-"`
}
