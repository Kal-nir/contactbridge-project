package model

type CustomerName struct {
	CustomerID        int    `json:"customer_id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	CustomerFirstName string `json:"customer_first_name"`
	CustomerSurname   string `json:"customer_surname"`
}
