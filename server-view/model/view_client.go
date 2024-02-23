package model

type ViewClient struct {
	NameID       int    `json:"name_id"`
	CustomerName string `json:"customer_name"`
	CompanyName  string `json:"company_name"`
}
