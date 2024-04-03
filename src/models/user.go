package models

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"not null unique"`
	Password string `json:"password" gorm:"not null"`
}
