package http

import (
	"svc-llt-golang/utils/logger"
	"svc-llt-golang/utils/response"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type HealthHandler struct {
	db *gorm.DB
}

func NewHealthHandler(db *gorm.DB) *HealthHandler {
	return &HealthHandler{db: db}
}

func (handler *HealthHandler) HealthCheck(ctx *fiber.Ctx) error {
	sql, err := handler.db.DB()
	if err != nil {
		logger.Error("Database connection error: " + err.Error())
		return response.Error(ctx, "Database connection error")
	}

	if err := sql.Ping(); err != nil {
		logger.Error("Database not reachable: " + err.Error())
		return response.InternalServerError(ctx, "Database not reachable")
	}

	logger.Info("Health check passed")
	return response.Success(ctx, "Service is running")
}