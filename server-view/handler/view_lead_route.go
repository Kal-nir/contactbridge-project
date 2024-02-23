package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kal-nir/contactbridge/server/database"
	"github.com/kal-nir/contactbridge/server/model"
)

func GetAllViewLeads(c *fiber.Ctx) error {
	db := database.DB.Db
	var view_leads []model.ViewLead

	db.Find(&view_leads)

	if len(view_leads) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "ViewLead not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "ViewLeads found",
		"data":    view_leads,
	})
}

func GetSingleViewLead(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")
	var view_lead model.ViewLead

	db.Find(&view_lead, "lead_id = ?", id)

	if view_lead.LeadID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "ViewLead not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "ViewLead found",
		"data":    view_lead,
	})
}
