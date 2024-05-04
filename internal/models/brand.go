package models

type Brand struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"not null" json:"name"`
	Country     string `gorm:"not null" json:"country"`
	Description string `gorm:"not null" json:"description"`
}
