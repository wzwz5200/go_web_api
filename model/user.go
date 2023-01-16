package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Names string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(110);not null"`
	Password string `gorm:"varchar(256);not null"`
}