package config

import (
	"fmt"
	"hospital-portal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func InitDB(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s port=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Kolkata",
		config.Host, config.User, config.Port, config.Password, config.DBName, config.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.User{}, &models.Patient{}); err != nil {
		return nil, fmt.Errorf("migration failed: %w", err)
	}
	return db, nil
}
