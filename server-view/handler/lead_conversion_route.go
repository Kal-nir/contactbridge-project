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
			"message": "Error with you input",
			"data":    err,
		})
	}

	err = db.Create(&lead_conversion).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Coud not create LeadConversion",
			"data":    err,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "LeadConversion has created",
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
		"message": "LeadConversions found",
		"data":    lead_conversions,
	})
}

func GetSingleLeadConversion(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")
	var lead_conversion model.LeadConversion

	db.Find(&lead_conversion, "lead_id = ?", id)

	if lead_conversion.LeadID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "LeadConversion not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "LeadConversion found",
		"data":    lead_conversion,
	})
}

func UpdateLeadConversion(c *fiber.Ctx) error {
	type updateLeadConversion struct {
		ConversionStatus  string `json:"conversion_status"`
		ConversionSource  string `json:"conversion_source"`
		ConversionRemarks string `json:"conversion_remarks"`
	}

	db := database.DB.Db
	var lead_conversion model.LeadConversion
	id := c.Params("id")

	db.Find(&lead_conversion, "lead_id = ?", id)

	if lead_conversion.LeadID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "User not found",
			"data":    nil,
		})
	}

	var updateLeadConversionData updateLeadConversion
	err := c.BodyParser(&updateLeadConversionData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Error with your input",
			"data":    err,
		})
	}

	lead_conversion.ConversionStatus = updateLeadConversionData.ConversionStatus
	lead_conversion.ConversionSource = updateLeadConversionData.ConversionSource
	lead_conversion.ConversionRemarks = updateLeadConversionData.ConversionRemarks

	db.Save(&lead_conversion)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "LeadConversion found",
		"data":    lead_conversion,
	})
}

func DeleteLeadConversionByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var lead_conversion model.LeadConversion
	id := c.Params("id")

	db.Find(&lead_conversion, "lead_id = ?", id)

	if lead_conversion.LeadID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "LeadConversion not found",
			"data":    nil,
		})
	}

	err := db.Delete(&lead_conversion, "lead_id = ?", id).Error

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
