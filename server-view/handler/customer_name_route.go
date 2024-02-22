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
			"message": "Something's wrong with your input",
			"data":    err,
		})
	}

	err = db.Create(&customer_name).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not create CustomerName",
			"data":    err,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "User has CustomerName",
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
		"message": "CustomerNames Found",
		"data":    customer_names,
	})
}

func GetSingleCustomerName(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")

	var customer_name model.CustomerName
	db.Find(&customer_name, "id = ?", id)

	if customer_name.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "CustomerName not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "CustomerName Found",
		"data":    customer_name,
	})
}

func UpdateCustomerName(c *fiber.Ctx) error {
	type updateCustomerName struct {
		FirstName string `json:"first_name"`
		Surname   string `json:"surname"`
	}

	db := database.DB.Db

	var customer_name model.CustomerName

	id := c.Params("id")

	db.Find(&customer_name, "id = ?", id)

	if customer_name.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "CustomerName not found",
			"data":    nil,
		})
	}

	var update_customer_name_data updateCustomerName
	err := c.BodyParser(&update_customer_name_data)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Something's wrong with your input",
			"data":    err,
		})
	}

	customer_name.FirstName = update_customer_name_data.FirstName
	customer_name.Surname = update_customer_name_data.Surname

	db.Save(&customer_name)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "CustomerName Found",
		"data":    customer_name,
	})
}

func DeleteCustomerNameByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var customer_name model.CustomerName

	id := c.Params("id")

	db.Find(&customer_name, "id = ?", id)

	if customer_name.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "CustomerName not found",
			"data":    nil,
		})
	}

	err := db.Delete(&customer_name, "id = ?", id).Error

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
