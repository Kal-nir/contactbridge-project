package model

type ViewLead struct {
	LeadID            int    `json:"lead_id"`
	ClientName        string `json:"client_name"`
	ClientCompany     string `json:"client_company"`
	ConversionStatus  string `json:"conversion_status"`
	ConversionSource  string `json:"conversion_source"`
	ConversionRemarks string `json:"conversion_remarks"`
}
