package models

import "time"

type Promotion struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"not null" json:"name"`
	Code           string    `gorm:"unique;not null" json:"code"`
	ExpirationDate time.Time `gorm:"not null" json:"expiration_date"`
}
