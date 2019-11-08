package migrations

import (
	"go-chi-gorm-test/database"
	"go-chi-gorm-test/models"
)

func Migrate() {

	database.DBCon.AutoMigrate(models.User{})

	// Seed
	// database.DBCon.Create(&models.User{ID: 1, FirstName: "HC", SecondName: "Goh"})
}
