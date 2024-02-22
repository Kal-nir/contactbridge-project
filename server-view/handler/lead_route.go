package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kal-nir/contactbridge/server/database"
	"github.com/kal-nir/contactbridge/server/model"
)

func GetAllLeads(c *fiber.Ctx) error {
	db := database.DB.Db
	var leads []model.Lead

	db.Find(&leads)

	if len(leads) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Lead not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Leads Found",
		"data":    leads,
	})
}

func GetSingleLead(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")

	var lead model.Lead
	db.Find(&lead, "id = ?", id)

	if lead.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Lead not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Lead Found",
		"data":    lead,
	})
}
