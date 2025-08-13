package v1

import (
	"svc-llt-golang/internal/middleware"
	userHttp "svc-llt-golang/internal/users/delivery/http"
	"svc-llt-golang/internal/users/repository"
	"svc-llt-golang/internal/users/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB, jwtKey string) {
	userRepo := repository.NewUserRepository(db)
	userUC := usecase.NewUserUsecase(userRepo, jwtKey)
	userHandler := userHttp.NewUserHandler(userUC)
	healthHandler := userHttp.NewHealthHandler(db)

	// Authentication routes
	auth := api.Group("/auth")
	auth.Post("/login", userHandler.Login)

	// Public routes
	// public := api.Group("/public/api/v1", middleware.JWTProtected())

	// Private routes
	private := api.Group("/private/api/v1", middleware.InternalApiKeyAuth())
	private.Get("/healthcheck", healthHandler.HealthCheck)
	private.Get("/healtchecker", healthHandler.HealthCheck)
}
