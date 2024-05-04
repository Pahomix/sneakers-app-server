package models

type Sneaker struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	Model    string  `gorm:"not null" json:"model"`
	BrandID  uint    `gorm:"not null" json:"brand_id"`
	Size     string  `gorm:"not null" json:"size"`
	Price    float64 `gorm:"not null" json:"price"`
	Quantity uint    `gorm:"not null" json:"quantity"`
}
