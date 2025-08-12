package v1

import (
	"svc-itasset-golang/internal/handler"
	"svc-itasset-golang/internal/middleware"
	"svc-itasset-golang/internal/repository"
	"svc-itasset-golang/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB, jwtKey string) {
	userRepo := repository.NewUserRepository(db)
	userUC := usecase.NewUserUsecase(userRepo, jwtKey)
	userHandler := handler.NewUserHandler(userUC)
	healthHandler := handler.NewHealthHandler(db)

	// Authentication routes
	auth := api.Group("/auth")
	auth.Post("/login", userHandler.Login)

	// Public routes
	// public := api.Group("/public/api/v1", middleware.JWTProtected())

	// Private routes
	private := api.Group("/private/api/v1", middleware.InternalApiKeyAuth())
	private.Get("/healthcheck", healthHandler.HealthCheck)
}
