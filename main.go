package main

import (
	"encoding/json"
	"fmt"
	"food-delivery/database"
	"food-delivery/migrations"
	"food-delivery/models"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	port := os.Getenv("PORT")
	fmt.Println(port)
	if port == "" {
		port = "8000"
	}

	database.InitDB()
	defer database.DBCon.Close()
	migrations.Migrate()

	r := chi.NewRouter()
	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		var users []models.User

		allusers := database.DBCon.Find(&users)
		json.NewEncoder(w).Encode(allusers)
	})

	r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		// request := map[string]string{}

		// json.NewDecoder(r.Body).Decode(&request)

		// fmt.Println(request.Test)
		r.ParseForm()

		fmt.Printf(r.FormValue("FirstName"))

		database.DBCon.Create(&models.User{ID: 6, FirstName: r.FormValue("FirstName"), SecondName: r.FormValue("LastName")})
	})

	http.ListenAndServe(":"+port, r)
}
