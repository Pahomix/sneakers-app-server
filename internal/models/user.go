package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Nickname string `gorm:"not null" json:"nickname"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`
	Role     string `gorm:"not null" json:"role"`
}
