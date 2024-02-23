package model

type ViewLead struct {
	LeadID            int    `json:"lead_id"`
	CustomerName      string `json:"customer_name"`
	CompanyName       string `json:"compnay_name"`
	ConversionStatus  string `json:"conversion_status"`
	ConversionSource  string `json:"conversion_source"`
	ConversionRemarks string `json:"conversion_remarks"`
}
