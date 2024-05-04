package models

type Category struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
}
