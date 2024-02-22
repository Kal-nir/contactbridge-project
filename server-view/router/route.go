package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kal-nir/contactbridge/server/handler"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	cnr := api.Group("/customer_name")
	cnr.Get("/", handler.GetAllCustomerNames)
	cnr.Get("/:id", handler.GetSingleCustomerName)
	cnr.Post("/", handler.CreateCustomerName)
	cnr.Put("/:id", handler.UpdateCustomerName)
	cnr.Delete("/:id", handler.DeleteCustomerNameByID)

	cmr := api.Group("/company_name")
	cmr.Get("/", handler.GetAllCompanyNames)
	cmr.Get("/:id", handler.GetSingleCompanyName)
	cmr.Post("/", handler.CreateCompanyName)
	cmr.Put("/:id", handler.UpdateCompanyName)
	cmr.Delete("/:id", handler.DeleteCompanyNameByID)

	ner := api.Group("/name_entity")
	ner.Get("/", handler.GetAllNameEntities)
	ner.Get("/:id", handler.GetSingleNameEntity)
	ner.Post("/", handler.CreateNameEntity)
	ner.Put("/:id", handler.UpdateNameEntity)
	ner.Delete("/:id", handler.DeleteNameEntityByID)

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
