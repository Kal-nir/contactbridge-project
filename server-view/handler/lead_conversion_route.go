package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kal-nir/contactbridge/server/database"
	"github.com/kal-nir/contactbridge/server/model"
)

func CreateLeadConversion(c *fiber.Ctx) error {
	db := database.DB.Db
	lead_conversion := new(model.LeadConversion)

	err := c.BodyParser(lead_conversion)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Something's wrong with your input",
			"data":    err,
		})
	}

	err = db.Create(&lead_conversion).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not create LeadConversion",
			"data":    err,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "User has LeadConversion",
		"data":    lead_conversion,
	})
}

func GetAllLeadConversions(c *fiber.Ctx) error {
	db := database.DB.Db
	var lead_conversions []model.LeadConversion

	db.Find(&lead_conversions)

	if len(lead_conversions) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "LeadConversion not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "LeadConversions Found",
		"data":    lead_conversions,
	})
}

func GetSingleLeadConversion(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")

	var lead_conversion model.LeadConversion
	db.Find(&lead_conversion, "id = ?", id)

	if lead_conversion.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "LeadConversion not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "LeadConversion Found",
		"data":    lead_conversion,
	})
}

func UpdateLeadConversion(c *fiber.Ctx) error {
	type updateLeadConversion struct {
		Status  string `json:"status"`
		Source  string `json:"source"`
		Remarks string `json:"remarks"`
	}

	db := database.DB.Db

	var lead_conversion model.LeadConversion

	id := c.Params("id")

	db.Find(&lead_conversion, "id = ?", id)

	if lead_conversion.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "LeadConversion not found",
			"data":    nil,
		})
	}

	var update_lead_conversion_data updateLeadConversion
	err := c.BodyParser(&update_lead_conversion_data)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Something's wrong with your input",
			"data":    err,
		})
	}

	lead_conversion.Status = update_lead_conversion_data.Status
	lead_conversion.Source = update_lead_conversion_data.Source
	lead_conversion.Remarks = update_lead_conversion_data.Remarks

	db.Save(&lead_conversion)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "LeadConversion Found",
		"data":    lead_conversion,
	})
}

func DeleteLeadConversionByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var lead_conversion model.LeadConversion

	id := c.Params("id")

	db.Find(&lead_conversion, "id = ?", id)

	if lead_conversion.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "LeadConversion not found",
			"data":    nil,
		})
	}

	err := db.Delete(&lead_conversion, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to delete LeadConversion",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "LeadConversion deleted",
	})
}
