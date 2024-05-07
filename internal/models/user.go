package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string  `gorm:"not null" json:"username"`
	Email     string  `gorm:"unique;not null" json:"email"`
	Password  string  `gorm:"not null" json:"-"`
	Role      string  `gorm:"not null" json:"role"`
	UserPhoto *string `gorm:"-" json:"user_photo,omitempty"`
}
