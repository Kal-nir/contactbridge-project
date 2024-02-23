package model

type LeadConversion struct {
	LeadID            int    `json:"lead_id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	ConversionStatus  string `json:"conversion_status"`
	ConversionSource  string `json:"conversion_source"`
	ConversionRemarks string `json:"conversion_remarks"`
}
