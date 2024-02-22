package model

import "gorm.io/gorm"

type CompanyName struct {
	gorm.Model
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Industry string `json:"industry"`
}

type CompanyNames struct {
	CompanyNames []CompanyName `json:"CompanyNames"`
}
