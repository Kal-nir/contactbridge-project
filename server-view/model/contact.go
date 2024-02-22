package model

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	ID           int    `json:"id"`
	CustomerName string `json:"customer_name"`
	CompanyName  string `json:"company_name"`
	LeadID       int    `json:"lead_id"`
	EmailAddress string `json:"email_address"`
	PhoneNumber  int    `json:"phone_number"`
	Note         string `json:"note"`
}

type Contacts struct {
	Contacts []Contact `json:"contacts"`
}
