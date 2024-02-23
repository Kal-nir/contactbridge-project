package model

type ViewContact struct {
	ClientID           int    `json:"client_id"`
	CustomerName       string `json:"customer_name"`
	CompanyName        string `json:"company_name"`
	LeadID             int    `json:"lead_id"`
	ClientEmailAddress string `json:"client_email_address"`
	ClientPhoneNumber  string `json:"client_phone_number"`
	ClientNote         string `json:"client_note"`
}
