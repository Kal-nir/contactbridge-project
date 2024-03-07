package model

type ClientContact struct {
	ClientID           int    `json:"client_id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	LeadID             int    `json:"lead_id" gorm:"default:null"`
	ClientEmailAddress string `json:"client_email_address"`
	ClientPhoneNumber  string `json:"client_phone_number"`
	ClientNote         string `json:"client_note"`
	ClientName         string `json:"client_name"`
	ClientCompany      string `json:"client_company"`
}
