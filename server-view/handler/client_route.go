package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kal-nir/contactbridge/server/database"
	"github.com/kal-nir/contactbridge/server/model"
)

func GetAllClients(c *fiber.Ctx) error {
	db := database.DB.Db
	var clients []model.Client

	db.Find(&clients)

	if len(clients) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Client not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Clients Found",
		"data":    clients,
	})
}

func GetSingleClient(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")

	var client model.Client
	db.Find(&client, "id = ?", id)

	if client.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Client not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Client Found",
		"data":    client,
	})
}
