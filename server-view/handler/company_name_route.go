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
			"message": "Something's wrong with your input",
			"data":    err,
		})
	}

	err = db.Create(&company_name).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not create CompanyName",
			"data":    err,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "User has CompanyName",
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
		"message": "CompanyNames Found",
		"data":    company_names,
	})
}

func GetSingleCompanyName(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")

	var company_name model.CompanyName
	db.Find(&company_name, "id = ?", id)

	if company_name.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "CompanyName not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "CompanyName Found",
		"data":    company_name,
	})
}

func UpdateCompanyName(c *fiber.Ctx) error {
	type updateCompanyName struct {
		Name     string `json:"name"`
		Industry string `json:"industry"`
	}

	db := database.DB.Db

	var company_name model.CompanyName

	id := c.Params("id")

	db.Find(&company_name, "id = ?", id)

	if company_name.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "CompanyName not found",
			"data":    nil,
		})
	}

	var update_company_name_data updateCompanyName
	err := c.BodyParser(&update_company_name_data)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Something's wrong with your input",
			"data":    err,
		})
	}

	company_name.Name = update_company_name_data.Name
	company_name.Industry = update_company_name_data.Industry

	db.Save(&company_name)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "CompanyName Found",
		"data":    company_name,
	})
}

func DeleteCompanyNameByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var company_name model.CompanyName

	id := c.Params("id")

	db.Find(&company_name, "id = ?", id)

	if company_name.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "CompanyName not found",
			"data":    nil,
		})
	}

	err := db.Delete(&company_name, "id = ?", id).Error

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
