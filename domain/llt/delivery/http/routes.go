package http

import (
	"svc-llt-golang/domain/llt"

	"github.com/gofiber/fiber/v2"
)

func NewLltRoutes(app *fiber.App, lltUsecase llt.Usecase) {
	// Initialize handlers
	lansiaHandler := NewLansiaHandler(lltUsecase)

	api := app.Group("/api/v1")
	lansia := api.Group("/lansia")

	{
		lansia.Get("/", lansiaHandler.GetAllLansia)
		lansia.Get("/detail", lansiaHandler.GetOneLansia)
		lansia.Post("/", lansiaHandler.StoreLansia)
		lansia.Put("/", lansiaHandler.UpdateLansia)
		lansia.Delete("/", lansiaHandler.DeleteLansia)
	}
}
