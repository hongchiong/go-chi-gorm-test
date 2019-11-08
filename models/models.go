package models

import "time"

type User struct {
	//this represents the customer
	ID int `json:"id,primary_key"` //this will be the primary key field

	FirstName  string `gorm:"size:255"`
	SecondName string `gorm:"size:255"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
