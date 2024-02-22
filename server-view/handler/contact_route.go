package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kal-nir/contactbridge/server/database"
	"github.com/kal-nir/contactbridge/server/model"
)

func GetAllContacts(c *fiber.Ctx) error {
	db := database.DB.Db
	var contacts []model.Contact

	db.Find(&contacts)

	if len(contacts) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Contact not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Contacts Found",
		"data":    contacts,
	})
}

func GetSingleContact(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")

	var contact model.Contact
	db.Find(&contact, "id = ?", id)

	if contact.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Contact not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Contact Found",
		"data":    contact,
	})
}
