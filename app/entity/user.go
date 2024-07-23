package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `gorm:"type:varchar(255);not null"`
}
