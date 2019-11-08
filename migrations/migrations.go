package migrations

import (
	"food-delivery/database"
	"food-delivery/models"
)

func Migrate() {

	database.DBCon.AutoMigrate(models.User{})

	// Seed
	// database.DBCon.Create(&models.User{ID: 1, FirstName: "HC", SecondName: "Goh"})
}
