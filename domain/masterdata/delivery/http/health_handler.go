package http

import (
	"svc-llt-golang/domain/masterdata"
	"svc-llt-golang/utils/logger"
	"svc-llt-golang/utils/response"

	"github.com/gofiber/fiber/v2"
)

type HealthHandler struct {
	masterdataUsecase masterdata.Usecase
}

func NewHealthHandler(masterdataUC masterdata.Usecase) *HealthHandler {
	return &HealthHandler{masterdataUsecase: masterdataUC}
}

func (handler *HealthHandler) HealthCheck(ctx *fiber.Ctx) error {
	if err := handler.masterdataUsecase.HealthCheck(); err != nil {
		logger.Error("Health check failed: " + err.Error())
		return response.InternalServerError(ctx, "Service unhealthy")
	}

	logger.Info("Health check passed")
	return response.Success(ctx, "Service is running")
}