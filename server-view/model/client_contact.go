package model

import "gorm.io/gorm"

type ClientContact struct {
	gorm.Model
	ID           int    `json:"id"`
	NameID       int    `json:"name_id"`
	LeadID       int    `json:"lead_id"`
	EmailAddress string `json:"email_address"`
	PhoneNumber  int    `json:"phone_number"`
	Note         string `json:"note"`
}

type ClientContacts struct {
	ClientContacts []ClientContact `json:"client_contacts"`
}
