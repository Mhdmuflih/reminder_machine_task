// package db

// import (
// 	"gorm.io/gorm"
// 	"gorm.io/driver/postgres"
// 	"os"
// 	"reminder/models"
// )

// var DB *gorm.DB

// func Connect() {
// 	dsn := os.Getenv("DB_URL")
// 	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("DB connection failed")
// 	}

// 	DB = database

// 	// AUTO CREATE TABLE
// 	DB.AutoMigrate(&models.User{})
// }

package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"reminder/models"
)

var DB *gorm.DB

func ConnectDB() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get env variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate tables
	err = db.AutoMigrate(
		&models.Task{},
		&models.ReminderRule{},
		&models.AuditLog{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	DB = db

	fmt.Println("Database connected successfully!")
}
