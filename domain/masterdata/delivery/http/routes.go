package http

import (
	"svc-llt-golang/domain/masterdata/repository"
	"svc-llt-golang/domain/masterdata/usecase"
	"svc-llt-golang/utils/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB, jwtKey string) {
	masterdataRepo := repository.NewMysqlMasterdataRepository(db)
	masterdataUC := usecase.NewMasterdataUsecase(masterdataRepo, jwtKey)
	userHandler := NewUserHandler(masterdataUC)
	healthHandler := NewHealthHandler(masterdataUC)

	// Authentication routes
	auth := api.Group("/auth")
	auth.Post("/login", userHandler.Login)

	// Private routes
	private := api.Group("/private/api/v1", middleware.InternalApiKeyAuth())
	private.Get("/healthcheck", healthHandler.HealthCheck)
	private.Get("/healtchecker", healthHandler.HealthCheck)
}