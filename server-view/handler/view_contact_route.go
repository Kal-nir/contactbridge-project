package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kal-nir/contactbridge/server/database"
	"github.com/kal-nir/contactbridge/server/model"
)

func GetAllViewContacts(c *fiber.Ctx) error {
	db := database.DB.Db
	var view_contacts []model.ViewContact

	db.Find(&view_contacts)

	if len(view_contacts) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "ViewContact not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "ViewContacts found",
		"data":    view_contacts,
	})
}

func GetSingleViewContact(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")
	var view_contact model.ViewContact

	db.Find(&view_contact, "client_id = ?", id)

	if view_contact.ClientID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "ViewContact not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "ViewContact found",
		"data":    view_contact,
	})
}
