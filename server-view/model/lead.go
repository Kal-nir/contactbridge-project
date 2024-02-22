package model

import "gorm.io/gorm"

type Lead struct {
	gorm.Model
	ID           int    `json:"id"`
	CustomerName string `json:"customer_name"`
	CompanyName  string `json:"company_name"`
	Status       string `json:"status"`
	Source       string `json:"source"`
	Remarks      string `json:"remarks"`
}

type Leads struct {
	Leads []Lead `json:"leads"`
}
