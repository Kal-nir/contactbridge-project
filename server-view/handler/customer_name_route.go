package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kal-nir/contactbridge/server/database"
	"github.com/kal-nir/contactbridge/server/model"
)

func CreateCustomerName(c *fiber.Ctx) error {
	db := database.DB.Db
	customer_name := new(model.CustomerName)

	err := c.BodyParser(customer_name)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Error with you input",
			"data":    err,
		})
	}

	err = db.Create(&customer_name).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Coud not create CustomerName",
			"data":    err,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "CustomerName has created",
		"data":    customer_name,
	})
}

func GetAllCustomerNames(c *fiber.Ctx) error {
	db := database.DB.Db
	var customer_names []model.CustomerName

	db.Find(&customer_names)

	if len(customer_names) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "CustomerName not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "CustomerNames found",
		"data":    customer_names,
	})
}

func GetSingleCustomerName(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")
	var customer_name model.CustomerName

	db.Find(&customer_name, "customer_id = ?", id)

	if customer_name.CustomerID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "CustomerName not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "CustomerName found",
		"data":    customer_name,
	})
}

func UpdateCustomerName(c *fiber.Ctx) error {
	type updateCustomerName struct {
		CustomerFirstName string `json:"customer_first_name"`
		CustomerSurname   string `json:"customer_surname"`
	}

	db := database.DB.Db
	var customer_name model.CustomerName
	id := c.Params("id")

	db.Find(&customer_name, "customer_id = ?", id)

	if customer_name.CustomerID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "User not found",
			"data":    nil,
		})
	}

	var updateCustomerNameData updateCustomerName
	err := c.BodyParser(&updateCustomerNameData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Error with your input",
			"data":    err,
		})
	}

	customer_name.CustomerFirstName = updateCustomerNameData.CustomerFirstName
	customer_name.CustomerSurname = updateCustomerNameData.CustomerSurname

	db.Save(&customer_name)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "CustomerName found",
		"data":    customer_name,
	})
}

func DeleteCustomerNameByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var customer_name model.CustomerName
	id := c.Params("id")

	db.Find(&customer_name, "customer_id = ?", id)

	if customer_name.CustomerID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "CustomerName not found",
			"data":    nil,
		})
	}

	err := db.Delete(&customer_name, "customer_id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to delete CustomerName",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "CustomerName deleted",
	})
}
