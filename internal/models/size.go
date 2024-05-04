package models

type Size struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	SneakerID uint   `gorm:"not null" json:"sneaker_id"`
	Size      string `gorm:"not null" json:"size"`
	Quantity  uint   `gorm:"not null" json:"quantity"`
}
