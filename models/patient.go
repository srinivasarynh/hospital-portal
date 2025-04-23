package models

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	FullName     string    `gorm:"not null" json:"full_name"`
	Age          int       `gorm:"not null;check:age > 0" json:"age"`
	Gender       string    `gorm:"type:VARCHAR(10);not null;check:gender IN ('Male','Female','Other')" json:"gender"`
	Symptoms     *string   `json:"symptoms"`
	Notes        string    `json:"notes"`
	RegisteredAt time.Time `gorm:"autoCreateTime" json:"registered_at"`
}
