package http

import (
	"log"
	"svc-llt-golang/domain/llt/repository"
	"svc-llt-golang/domain/llt/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterLltRoutes(api fiber.Router, db *gorm.DB) {
	// Initialize repository and usecase
	lltRepo := repository.NewMysqlLltRepository(db)
	lltUC := usecase.NewLltUsecase(lltRepo)
	lansiaHandler := NewLansiaHandler(lltUC)

	// Lansia routes under /lansia
	lansia := api.Group("/lansia")
	{
		lansia.Get("/", lansiaHandler.GetAllLansia)
		lansia.Get("/detail", lansiaHandler.GetOneLansia)
		lansia.Post("/", lansiaHandler.StoreLansia)
		lansia.Put("/", lansiaHandler.UpdateLansia)
		lansia.Delete("/", lansiaHandler.DeleteLansia)
	}
	
	// Debug: Log that routes are registered
	log.Println("LLT routes registered successfully")
}
