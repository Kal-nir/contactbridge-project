package model

import "gorm.io/gorm"

type NameEntity struct {
	gorm.Model
	ID         int `json:"id"`
	CustomerID int `json:"customer_id"`
	CompanyID  int `json:"company_id"`
}

type NameEntities struct {
	NameEntities []NameEntity `json:"name_entities"`
}
