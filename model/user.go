package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Names     string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(110);not null"`
	Password  string `gorm:"varchar(256);not null"`
	EndTime   time.Time
	VIP_TIME  time.Time
}

type Building struct {
	gorm.Model
	EndTime time.Time
}
