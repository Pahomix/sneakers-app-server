package models

import "gorm.io/gorm"

type Sneaker struct {
	ID         uint            `gorm:"primaryKey" json:"id"`
	Model      string          `gorm:"not null" json:"model"`
	CategoryID uint            `gorm:"not null" json:"category_id"`
	BrandID    uint            `gorm:"not null" json:"brand_id"`
	Sizes      []Size          `gorm:"many2many:sneaker_sizes;not null" json:"sizes"`
	Price      float64         `gorm:"not null" json:"price"`
	Quantity   uint            `gorm:"not null" json:"quantity"`
	Photo      *string         `gorm:"-" json:"photo,omitempty"`
	DeletedAt  *gorm.DeletedAt `gorm:"index" json:"-"`
}
