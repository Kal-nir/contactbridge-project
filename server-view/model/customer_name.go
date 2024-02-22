package model

import "gorm.io/gorm"

type CustomerName struct {
	gorm.Model
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	Surname   string `json:"surname"`
}

type CustomerNames struct {
	CustomerNames []CustomerName `json:"customer_names"`
}
