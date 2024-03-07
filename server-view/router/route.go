package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kal-nir/contactbridge/server/handler"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	vlr := api.Group("/view_lead")
	vlr.Get("/", handler.GetAllViewLeads)
	vlr.Get("/:id", handler.GetSingleViewLead)

	ccr := api.Group("/client_contact")
	ccr.Get("/", handler.GetAllClientContacts)
	ccr.Get("/:id", handler.GetSingleClientContact)
	ccr.Post("/", handler.CreateClientContact)
	ccr.Put("/:id", handler.UpdateClientContact)
	ccr.Delete("/:id", handler.DeleteClientContactByID)

	lcr := api.Group("/lead_conversion")
	lcr.Get("/", handler.GetAllLeadConversions)
	lcr.Get("/:id", handler.GetSingleLeadConversion)
	lcr.Post("/", handler.CreateLeadConversion)
	lcr.Put("/:id", handler.UpdateLeadConversion)
	lcr.Delete("/:id", handler.DeleteLeadConversionByID)
}
