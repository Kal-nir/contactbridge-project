package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kal-nir/contactbridge/server/database"
	"github.com/kal-nir/contactbridge/server/model"
)

func CreateClientContact(c *fiber.Ctx) error {
	db := database.DB.Db
	client_contact := new(model.ClientContact)

	err := c.BodyParser(client_contact)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Error with your input",
			"data":    err.Error(),
		})
	}

	err = db.Create(&client_contact).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not create ClientContact",
			"data":    err,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "ClientContact has created",
		"data":    client_contact,
	})
}

func GetAllClientContacts(c *fiber.Ctx) error {
	db := database.DB.Db
	var client_contacts []model.ClientContact

	db.Find(&client_contacts)

	if len(client_contacts) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "ClientContact not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "ClientContacts found",
		"data":    client_contacts,
	})
}

func GetSingleClientContact(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")
	var client_contact model.ClientContact

	db.Find(&client_contact, "client_id = ?", id)

	if client_contact.ClientID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "ClientContact not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "ClientContact found",
		"data":    client_contact,
	})
}

func UpdateClientContact(c *fiber.Ctx) error {
	type updateClientContact struct {
		LeadID             int    `json:"lead_id" gorm:"default:null"`
		ClientEmailAddress string `json:"client_email_address"`
		ClientPhoneNumber  string `json:"client_phone_number"`
		ClientNote         string `json:"client_note"`
		ClientName         string `json:"client_name"`
		ClientCompany      string `json:"client_company"`
	}

	db := database.DB.Db
	var client_contact model.ClientContact
	id := c.Params("id")

	db.Find(&client_contact, "client_id = ?", id)

	if client_contact.ClientID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "User not found",
			"data":    nil,
		})
	}

	var updateClientContactData updateClientContact
	err := c.BodyParser(&updateClientContactData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Error with your input",
			"data":    err,
		})
	}

	client_contact.LeadID = updateClientContactData.LeadID
	client_contact.ClientEmailAddress = updateClientContactData.ClientEmailAddress
	client_contact.ClientPhoneNumber = updateClientContactData.ClientPhoneNumber
	client_contact.ClientNote = updateClientContactData.ClientNote
	client_contact.ClientName = updateClientContactData.ClientName
	client_contact.ClientCompany = updateClientContactData.ClientCompany

	lead := &model.LeadConversion{}
	var exists bool
	err = db.Model(lead).
		Select("count(*) > 0").
		Where("lead_id = ?", updateClientContactData.LeadID).
		Find(&exists).
		Error

	if !exists {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "The inputted lead ID must exist!",
			"data":    err,
		})
	}

	db.Save(&client_contact)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "ClientContact found",
		"data":    client_contact,
	})
}

func DeleteClientContactByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var company_name model.ClientContact
	id := c.Params("id")

	db.Find(&company_name, "client_id = ?", id)

	if company_name.ClientID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "ClientContact not found",
			"data":    nil,
		})
	}

	err := db.Delete(&company_name, "client_id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to delete ClientContact",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "ClientContact deleted",
	})
}
