package http

import (
	"svc-llt-golang/domain/masterdata/repository"
	"svc-llt-golang/domain/masterdata/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB, jwtKey string) {
	masterdataRepo := repository.NewMysqlMasterdataRepository(db)
	masterdataUC := usecase.NewMasterdataUsecase(masterdataRepo, jwtKey)
	userHandler := NewUserHandler(masterdataUC)

	// Authentication routes
	auth := api.Group("/auth")
	auth.Post("/login", userHandler.Login)
	auth.Post("/register", userHandler.Register)
}