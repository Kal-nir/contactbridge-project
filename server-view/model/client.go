package model

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	ID           int    `json:"id"`
	CustomerName string `json:"customer_name"`
	CompanyName  string `json:"company_name"`
}

type Clients struct {
	Clients []Client `json:"clients"`
}
