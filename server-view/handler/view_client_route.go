package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kal-nir/contactbridge/server/database"
	"github.com/kal-nir/contactbridge/server/model"
)

func GetAllViewClients(c *fiber.Ctx) error {
	db := database.DB.Db
	var view_clients []model.ViewClient

	db.Find(&view_clients)

	if len(view_clients) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "ViewClient not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "ViewClients found",
		"data":    view_clients,
	})
}

func GetSingleViewClient(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")
	var view_client model.ViewClient

	db.Find(&view_client, "name_id = ?", id)

	if view_client.NameID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "ViewClient not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "ViewClient found",
		"data":    view_client,
	})
}
