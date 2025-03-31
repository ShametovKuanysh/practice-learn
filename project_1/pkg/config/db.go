package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading ENV file: ", err)
		return err
	}

	dbURL := os.Getenv("DB_URL")

	DB, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
		return err
	}

	db = DB
	return nil
}

func GetDB() *gorm.DB {
	return db
}
