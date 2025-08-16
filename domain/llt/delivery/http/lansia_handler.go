package http

import (
	"log"
	"svc-llt-golang/domain/llt"
	"svc-llt-golang/utils/header"
	"svc-llt-golang/utils/logger"
	"svc-llt-golang/utils/response"
	"svc-llt-golang/valueobject"

	"github.com/gofiber/fiber/v2"
)

type LansiaHandler struct {
	lltUsecase llt.Usecase
}

func NewLansiaHandler(lltUC llt.Usecase) *LansiaHandler {
	return &LansiaHandler{lltUsecase: lltUC}
}

func (handler *LansiaHandler) GetAllLansia(ctx *fiber.Ctx) error {
	log.Println("Get all lansia request received")

	param := map[string]interface{}{}

	if status := ctx.Query("status"); status != "" {
		param["status"] = status
	}

	lansiaList, err := handler.lltUsecase.GetAllLansia(param)
	if err != nil {
		logger.Error("Failed to get lansia data: " + err.Error())
		return response.Error(ctx, "Failed to get lansia data")
	}

	logger.Info("Successfully retrieved lansia data")
	return response.Success(ctx, valueobject.LansiaListResponse{
		Message: "Successfully retrieved lansia data",
		Data:    lansiaList,
		Total:   len(lansiaList),
	})
}

func (handler *LansiaHandler) GetOneLansia(ctx *fiber.Ctx) error {
	log.Println("Get one lansia request received")

	uuid := ctx.Query("uuid")
	if uuid == "" {
		logger.Error("UUID query parameter is required")
		return response.BadRequest(ctx, "UUID query parameter is required")
	}

	param := map[string]interface{}{
		"uuid": uuid,
	}

	lansia, err := handler.lltUsecase.GetOneLansia(param)
	if err != nil {
		logger.Error("Failed to get lansia data: " + err.Error())
		return response.NotFound(ctx, "Lansia not found")
	}

	logger.Info("Successfully retrieved lansia data for UUID: " + uuid)
	return response.Success(ctx, lansia)
}

func (handler *LansiaHandler) StoreLansia(ctx *fiber.Ctx) error {
	log.Println("Store lansia request received")
	var req valueobject.LansiaPayloadInsert

	if err := ctx.BodyParser(&req); err != nil {
		logger.Error("Failed to parse request body: " + err.Error())
		return response.BadRequest(ctx, "Invalid request payload")
	}

	xMember, err := header.ValidateAndExtractXMember(ctx)
	if err != nil {
		logger.Error("X-Member header validation failed: " + err.Error())
		return err
	}
	req.User = xMember

	result, err := handler.lltUsecase.StoreLansia(req)
	if err != nil {
		logger.Error("Failed to store lansia: " + err.Error())
		return response.Error(ctx, err.Error())
	}

	logger.Info("Lansia stored successfully")
	return response.Success(ctx, result)
}

func (handler *LansiaHandler) UpdateLansia(ctx *fiber.Ctx) error {
	log.Println("Update lansia request received")
	var req valueobject.LansiaPayloadUpdate

	if err := ctx.BodyParser(&req); err != nil {
		logger.Error("Failed to parse request body: " + err.Error())
		return response.BadRequest(ctx, "Invalid request payload")
	}

	xMember, err := header.ValidateAndExtractXMember(ctx)
	if err != nil {
		logger.Error("X-Member header validation failed: " + err.Error())
		return err
	}
	req.User = xMember

	err = handler.lltUsecase.UpdateLansia(req)
	if err != nil {
		logger.Error("Failed to update lansia: " + err.Error())
		return response.Error(ctx, err.Error())
	}

	logger.Info("Lansia updated successfully")
	return response.Success(ctx, "Lansia updated successfully")
}

func (handler *LansiaHandler) DeleteLansia(ctx *fiber.Ctx) error {
	log.Println("Delete lansia request received")
	var req valueobject.LansiaPayloadDelete

	if err := ctx.BodyParser(&req); err != nil {
		logger.Error("Failed to parse request body: " + err.Error())
		return response.BadRequest(ctx, "Invalid request payload")
	}

	xMember, err := header.ValidateAndExtractXMember(ctx)
	if err != nil {
		logger.Error("X-Member header validation failed: " + err.Error())
		return err
	}
	req.User = xMember

	err = handler.lltUsecase.DeleteLansia(req)
	if err != nil {
		logger.Error("Failed to delete lansia: " + err.Error())
		return response.Error(ctx, err.Error())
	}

	logger.Info("Lansia deleted successfully")
	return response.Success(ctx, "Lansia deleted successfully")
}
