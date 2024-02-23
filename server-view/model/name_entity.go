package model

type NameEntity struct {
	NameID     int `json:"name_id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	CustomerID int `json:"customer_id"`
	CompanyID  int `json:"company_id"`
}
