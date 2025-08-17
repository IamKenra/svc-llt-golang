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

type IdentitasHandler struct {
	lltUsecase llt.Usecase
}

func NewIdentitasHandler(lltUC llt.Usecase) *IdentitasHandler {
	return &IdentitasHandler{lltUsecase: lltUC}
}

func (handler *IdentitasHandler) GetAllIdentitas(ctx *fiber.Ctx) error {
	log.Println("Get all identitas request received")

	param := map[string]interface{}{}

	if nik := ctx.Query("nik"); nik != "" {
		param["nik"] = nik
	}

	identitasList, err := handler.lltUsecase.GetAllIdentitas(param)
	if err != nil {
		logger.Error("Failed to get identitas data: " + err.Error())
		return response.Error(ctx, "Failed to get identitas data")
	}

	logger.Info("Successfully retrieved identitas data")
	return response.Success(ctx, valueobject.IdentitasListResponse{
		Message: "Successfully retrieved identitas data",
		Data:    identitasList,
		Total:   len(identitasList),
	})
}

func (handler *IdentitasHandler) GetOneIdentitas(ctx *fiber.Ctx) error {
	log.Println("Get one identitas request received")

	uuid := ctx.Query("uuid")
	if uuid == "" {
		logger.Error("UUID query parameter is required")
		return response.BadRequest(ctx, "UUID query parameter is required")
	}

	param := map[string]interface{}{
		"uuid": uuid,
	}

	identitas, err := handler.lltUsecase.GetOneIdentitas(param)
	if err != nil {
		logger.Error("Failed to get identitas data: " + err.Error())
		return response.NotFound(ctx, "Identitas not found")
	}

	logger.Info("Successfully retrieved identitas data for UUID: " + uuid)
	return response.Success(ctx, identitas)
}

func (handler *IdentitasHandler) StoreIdentitas(ctx *fiber.Ctx) error {
	log.Println("Store identitas request received")
	var req valueobject.IdentitasPayloadInsert

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

	result, err := handler.lltUsecase.StoreIdentitas(req)
	if err != nil {
		logger.Error("Failed to store identitas: " + err.Error())
		return response.Error(ctx, err.Error())
	}

	logger.Info("Identitas stored successfully")
	return response.Success(ctx, result)
}

func (handler *IdentitasHandler) UpdateIdentitas(ctx *fiber.Ctx) error {
	log.Println("Update identitas request received")
	var req valueobject.IdentitasPayloadUpdate

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

	err = handler.lltUsecase.UpdateIdentitas(req)
	if err != nil {
		logger.Error("Failed to update identitas: " + err.Error())
		return response.Error(ctx, err.Error())
	}

	logger.Info("Identitas updated successfully")
	return response.Success(ctx, "Identitas updated successfully")
}

func (handler *IdentitasHandler) DeleteIdentitas(ctx *fiber.Ctx) error {
	log.Println("Delete identitas request received")
	var req valueobject.IdentitasPayloadDelete

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

	err = handler.lltUsecase.DeleteIdentitas(req)
	if err != nil {
		logger.Error("Failed to delete identitas: " + err.Error())
		return response.Error(ctx, err.Error())
	}

	logger.Info("Identitas deleted successfully")
	return response.Success(ctx, "Identitas deleted successfully")
}