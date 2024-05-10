package db

import "sneakers-app/internal/models"

func SyncDatabase() {
	Db.AutoMigrate(&models.User{})
	Db.AutoMigrate(&models.Sneaker{})
	Db.AutoMigrate(&models.Brand{})
	//Db.AutoMigrate(&models.Size{})
	Db.AutoMigrate(&models.Promotion{})
	Db.AutoMigrate(&models.Review{})
	Db.AutoMigrate(&models.Order{})
	Db.AutoMigrate(&models.PurchaseHistory{})
	Db.AutoMigrate(&models.Category{})
	Db.AutoMigrate(&models.CartItem{})
}
