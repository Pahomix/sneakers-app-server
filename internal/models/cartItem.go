package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	UserID     uint      `gorm:"not null" json:"user_id"`
	Quantity   uint      `gorm:"not null" json:"quantity"`
	TotalPrice float64   `gorm:"not null" json:"total_price"`
	Sneakers   []Sneaker `gorm:"many2many:cart_item_sneakers;not null" json:"sneakers"`
}

func (c *CartItem) AfterSave(tx *gorm.DB) (err error) {
	var totalAmount float64
	tx.Model(&CartItem{}).Where("user_id = ?", c.UserID).Select("SUM(total_price)").Row().Scan(&totalAmount)

	var order Order
	tx.Model(&Order{}).Where("user_id = ? AND status = ?", c.UserID, "не оплачено").First(&order)
	tx.Model(&Order{}).Where("user_id = ? AND status = ?", c.UserID, "не оплачено").Updates(map[string]interface{}{"total_amount": totalAmount})

	return
}
