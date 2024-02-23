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
			"message": "Error with you input",
			"data":    err,
		})
	}

	err = db.Create(&name_entity).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Coud not create NameEntity",
			"data":    err,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "NameEntity has created",
		"data":    name_entity,
	})
}

func GetAllNameEntities(c *fiber.Ctx) error {
	db := database.DB.Db
	var name_entities []model.NameEntity

	db.Find(&name_entities)

	if len(name_entities) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "NameEntity not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "NameEntitys found",
		"data":    name_entities,
	})
}

func GetSingleNameEntity(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")
	var name_entity model.NameEntity

	db.Find(&name_entity, "name_id = ?", id)

	if name_entity.CompanyID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "NameEntity not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "NameEntity found",
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

	db.Find(&name_entity, "name_id = ?", id)

	if name_entity.CompanyID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "User not found",
			"data":    nil,
		})
	}

	var updateNameEntityData updateNameEntity
	err := c.BodyParser(&updateNameEntityData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Error with your input",
			"data":    err,
		})
	}

	name_entity.CustomerID = updateNameEntityData.CustomerID
	name_entity.CompanyID = updateNameEntityData.CompanyID

	db.Save(&name_entity)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "NameEntity found",
		"data":    name_entity,
	})
}

func DeleteNameEntityByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var name_entity model.NameEntity
	id := c.Params("id")

	db.Find(&name_entity, "name_id = ?", id)

	if name_entity.CustomerID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "NameEntity not found",
			"data":    nil,
		})
	}

	err := db.Delete(&name_entity, "name_id = ?", id).Error

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
