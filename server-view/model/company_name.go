package model

type CompanyName struct {
	CompanyID       int    `json:"company_id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	CompanyName     string `json:"company_name"`
	CompanyIndustry string `json:"company_industry"`
}
