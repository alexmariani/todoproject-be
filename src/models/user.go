package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `gorm:"primaryKey"`
	Username string `json:"username" gorm:"not null unique"`
	Password string `json:"password" gorm:"not null"`
}
