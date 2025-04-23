package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"hospital-portal/models"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	seedFlag = flag.Bool("seed", false, "Seed default doctor and receptionist users")
)

func main() {
	godotenv.Load()
	flag.Parse()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	password := os.Getenv("DB_PASS")
	user := os.Getenv("DB_USER")
	DBname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s port=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Kolkata",
		host, user, port, password, DBname,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if *seedFlag {
		fmt.Println("Seeding users...")
		createUser(db, "doctor1", "doctorpassword", "doctor")
		createUser(db, "receptionist1", "receptionistpassword", "receptionist")
		fmt.Println("Users seeded successfully.")
	}
}

func createUser(db *gorm.DB, username, password, role string) {
	var existing models.User
	if err := db.Where("username = ?", username).First(&existing).Error; err == nil {
		fmt.Printf("User '%s' already exists, skipping...\n", username)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash password:", err)
	}

	user := models.User{
		Username: username,
		Password: string(hashedPassword),
		Role:     role,
	}

	if err := db.Create(&user).Error; err != nil {
		log.Fatal("Error creating user:", err)
	}
}
