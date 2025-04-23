package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"type:VARCHAR(20);not null;check:role IN ('doctor','receptionist')"`
}
