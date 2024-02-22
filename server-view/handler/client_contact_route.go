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
			"message": "Something's wrong with your input",
			"data":    err,
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
		"message": "User has ClientContact",
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
		"message": "ClientContacts Found",
		"data":    client_contacts,
	})
}

func GetSingleClientContact(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")

	var client_contact model.ClientContact
	db.Find(&client_contact, "id = ?", id)

	if client_contact.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "ClientContact not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "ClientContact Found",
		"data":    client_contact,
	})
}

func UpdateClientContact(c *fiber.Ctx) error {
	type updateClientContact struct {
		NameID       int    `json:"name_id"`
		LeadID       int    `json:"lead_id"`
		EmailAddress string `json:"email_address"`
		PhoneNumber  int    `json:"phone_number"`
		Note         string `json:"note"`
	}

	db := database.DB.Db

	var client_contact model.ClientContact

	id := c.Params("id")

	db.Find(&client_contact, "id = ?", id)

	if client_contact.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "ClientContact not found",
			"data":    nil,
		})
	}

	var update_client_contact_data updateClientContact
	err := c.BodyParser(&update_client_contact_data)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Something's wrong with your input",
			"data":    err,
		})
	}

	client_contact.NameID = update_client_contact_data.NameID
	client_contact.LeadID = update_client_contact_data.LeadID
	client_contact.EmailAddress = update_client_contact_data.EmailAddress
	client_contact.PhoneNumber = update_client_contact_data.PhoneNumber
	client_contact.Note = update_client_contact_data.Note

	db.Save(&client_contact)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "ClientContact Found",
		"data":    client_contact,
	})
}

func DeleteClientContactByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var client_contact model.ClientContact

	id := c.Params("id")

	db.Find(&client_contact, "id = ?", id)

	if client_contact.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "ClientContact not found",
			"data":    nil,
		})
	}

	err := db.Delete(&client_contact, "id = ?", id).Error

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
