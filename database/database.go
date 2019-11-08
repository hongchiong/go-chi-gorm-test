package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var (
	DBCon *gorm.DB
)

func InitDB() {
	var err error

	envError := godotenv.Load()
	if envError != nil {
		log.Println("failed to load godotenv")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	dbConfig := fmt.Sprintf("host=%s port=%v user=%s dbname=%s password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)

	DBCon, err = gorm.Open("postgres", dbConfig)

	if err != nil {
		panic("failed to connect database")
	}

}
