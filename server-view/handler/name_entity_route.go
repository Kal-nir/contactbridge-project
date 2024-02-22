package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kal-nir/contactbridge/server/database"
	"github.com/kal-nir/contactbridge/server/model"
)

func CreateNameEntity(c *fiber.Ctx) error {
	db := database.DB.Db
	name_entity := new(model.NameEntity)

	err := c.BodyParser(name_entity)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Something's wrong with your input",
			"data":    err,
		})
	}

	err = db.Create(&name_entity).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not create NameEntity",
			"data":    err,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "User has NameEntity",
		"data":    name_entity,
	})
}

func GetAllNameEntities(c *fiber.Ctx) error {
	db := database.DB.Db
	var name_entitys []model.NameEntity

	db.Find(&name_entitys)

	if len(name_entitys) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "NameEntity not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "NameEntitys Found",
		"data":    name_entitys,
	})
}

func GetSingleNameEntity(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")

	var name_entity model.NameEntity
	db.Find(&name_entity, "id = ?", id)

	if name_entity.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "NameEntity not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "NameEntity Found",
		"data":    name_entity,
	})
}

func UpdateNameEntity(c *fiber.Ctx) error {
	type updateNameEntity struct {
		CustomerID int `json:"customer_id"`
		CompanyID  int `json:"company_id"`
	}

	db := database.DB.Db

	var name_entity model.NameEntity

	id := c.Params("id")

	db.Find(&name_entity, "id = ?", id)

	if name_entity.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "NameEntity not found",
			"data":    nil,
		})
	}

	var update_name_entity_data updateNameEntity
	err := c.BodyParser(&update_name_entity_data)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Something's wrong with your input",
			"data":    err,
		})
	}

	name_entity.CustomerID = update_name_entity_data.CustomerID
	name_entity.CompanyID = update_name_entity_data.CompanyID

	db.Save(&name_entity)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "NameEntity Found",
		"data":    name_entity,
	})
}

func DeleteNameEntityByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var name_entity model.NameEntity

	id := c.Params("id")

	db.Find(&name_entity, "id = ?", id)

	if name_entity.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "NameEntity not found",
			"data":    nil,
		})
	}

	err := db.Delete(&name_entity, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to delete NameEntity",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "NameEntity deleted",
	})
}
