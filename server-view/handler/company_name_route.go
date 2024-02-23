package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kal-nir/contactbridge/server/database"
	"github.com/kal-nir/contactbridge/server/model"
)

func CreateCompanyName(c *fiber.Ctx) error {
	db := database.DB.Db
	company_name := new(model.CompanyName)

	err := c.BodyParser(company_name)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Error with you input",
			"data":    err,
		})
	}

	err = db.Create(&company_name).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Coud not create CompanyName",
			"data":    err,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "CompanyName has created",
		"data":    company_name,
	})
}

func GetAllCompanyNames(c *fiber.Ctx) error {
	db := database.DB.Db
	var company_names []model.CompanyName

	db.Find(&company_names)

	if len(company_names) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "CompanyName not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "CompanyNames found",
		"data":    company_names,
	})
}

func GetSingleCompanyName(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")
	var company_name model.CompanyName

	db.Find(&company_name, "company_id = ?", id)

	if company_name.CompanyID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "CompanyName not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "CompanyName found",
		"data":    company_name,
	})
}

func UpdateCompanyName(c *fiber.Ctx) error {
	type updateCompanyName struct {
		CompanyName     string `json:"company_name"`
		CompanyIndustry string `json:"company_industry"`
	}

	db := database.DB.Db
	var company_name model.CompanyName
	id := c.Params("id")

	db.Find(&company_name, "company_id = ?", id)

	if company_name.CompanyID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "User not found",
			"data":    nil,
		})
	}

	var updateCompanyNameData updateCompanyName
	err := c.BodyParser(&updateCompanyNameData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Error with your input",
			"data":    err,
		})
	}

	company_name.CompanyName = updateCompanyNameData.CompanyName
	company_name.CompanyIndustry = updateCompanyNameData.CompanyIndustry

	db.Save(&company_name)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "CompanyName found",
		"data":    company_name,
	})
}

func DeleteCompanyNameByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var company_name model.CompanyName
	id := c.Params("id")

	db.Find(&company_name, "company_id = ?", id)

	if company_name.CompanyID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "CompanyName not found",
			"data":    nil,
		})
	}

	err := db.Delete(&company_name, "company_id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to delete CompanyName",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "CompanyName deleted",
	})
}
