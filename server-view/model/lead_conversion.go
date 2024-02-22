package model

import "gorm.io/gorm"

type LeadConversion struct {
	gorm.Model
	ID      int    `json:"id"`
	Status  string `json:"status"`
	Source  string `json:"source"`
	Remarks string `json:"remarks"`
}

type LeadConversions struct {
	LeadConversions []LeadConversion `json:"lead_conversions"`
}
