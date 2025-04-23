package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"hospital-portal/config"
	"hospital-portal/routes"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	DBconfig := &config.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := config.InitDB(DBconfig)
	if err != nil {
		log.Fatal("Database connection failed", err)
	}

	router := gin.Default()
	routes.SetupRoutes(router, db)
	router.Run()
}
