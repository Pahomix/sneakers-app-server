package models

import "gorm.io/gorm"

type Sneaker struct {
	gorm.Model
	Title      string  `gorm:"not null" json:"title"`
	CategoryID uint    `gorm:"not null" json:"category_id"`
	BrandID    uint    `gorm:"not null" json:"brand_id"`
	Sizes      []Size  `gorm:"foreignkey:SneakerID" json:"sizes"`
	Price      float64 `gorm:"not null" json:"price"`
	Quantity   uint    `gorm:"not null" json:"quantity"`
	Photo      *string `gorm:"-" json:"photo,omitempty"`
}
