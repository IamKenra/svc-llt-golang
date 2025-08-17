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

type AlamatHandler struct {
	lltUsecase llt.Usecase
}

func NewAlamatHandler(lltUC llt.Usecase) *AlamatHandler {
	return &AlamatHandler{lltUsecase: lltUC}
}

func (handler *AlamatHandler) GetAllAlamat(ctx *fiber.Ctx) error {
	log.Println("Get all alamat request received")

	param := map[string]interface{}{}

	if kota := ctx.Query("kota"); kota != "" {
		param["kota"] = kota
	}
	if kecamatan := ctx.Query("kecamatan"); kecamatan != "" {
		param["kecamatan"] = kecamatan
	}

	alamatList, err := handler.lltUsecase.GetAllAlamat(param)
	if err != nil {
		logger.Error("Failed to get alamat data: " + err.Error())
		return response.Error(ctx, "Failed to get alamat data")
	}

	logger.Info("Successfully retrieved alamat data")
	return response.Success(ctx, valueobject.AlamatListResponse{
		Message: "Successfully retrieved alamat data",
		Data:    alamatList,
		Total:   len(alamatList),
	})
}

func (handler *AlamatHandler) GetOneAlamat(ctx *fiber.Ctx) error {
	log.Println("Get one alamat request received")

	uuid := ctx.Query("uuid")
	if uuid == "" {
		logger.Error("UUID query parameter is required")
		return response.BadRequest(ctx, "UUID query parameter is required")
	}

	param := map[string]interface{}{
		"uuid": uuid,
	}

	alamat, err := handler.lltUsecase.GetOneAlamat(param)
	if err != nil {
		logger.Error("Failed to get alamat data: " + err.Error())
		return response.NotFound(ctx, "Alamat not found")
	}

	logger.Info("Successfully retrieved alamat data for UUID: " + uuid)
	return response.Success(ctx, alamat)
}

func (handler *AlamatHandler) StoreAlamat(ctx *fiber.Ctx) error {
	log.Println("Store alamat request received")
	var req valueobject.AlamatPayloadInsert

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

	result, err := handler.lltUsecase.StoreAlamat(req)
	if err != nil {
		logger.Error("Failed to store alamat: " + err.Error())
		return response.Error(ctx, err.Error())
	}

	logger.Info("Alamat stored successfully")
	return response.Success(ctx, result)
}

func (handler *AlamatHandler) UpdateAlamat(ctx *fiber.Ctx) error {
	log.Println("Update alamat request received")
	var req valueobject.AlamatPayloadUpdate

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

	err = handler.lltUsecase.UpdateAlamat(req)
	if err != nil {
		logger.Error("Failed to update alamat: " + err.Error())
		return response.Error(ctx, err.Error())
	}

	logger.Info("Alamat updated successfully")
	return response.Success(ctx, "Alamat updated successfully")
}

func (handler *AlamatHandler) DeleteAlamat(ctx *fiber.Ctx) error {
	log.Println("Delete alamat request received")
	var req valueobject.AlamatPayloadDelete

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

	err = handler.lltUsecase.DeleteAlamat(req)
	if err != nil {
		logger.Error("Failed to delete alamat: " + err.Error())
		return response.Error(ctx, err.Error())
	}

	logger.Info("Alamat deleted successfully")
	return response.Success(ctx, "Alamat deleted successfully")
}