package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading.env file")
	}

	dbUrl := os.Getenv("DB_URL")

	d, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
