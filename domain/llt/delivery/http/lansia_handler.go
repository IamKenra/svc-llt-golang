package http

import (
	"log"
	"strings"
	"svc-llt-golang/domain/llt"
	"svc-llt-golang/entity"
	"svc-llt-golang/utils/header"
	"svc-llt-golang/utils/image"
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

	xMember, err := header.ValidateAndExtractXMember(ctx)
	if err != nil {
		logger.Error("X-Member header validation failed: " + err.Error())
		return err
	}

	// Check if request contains multipart form data (for image upload)
	contentType := ctx.Get("Content-Type")
	if strings.Contains(contentType, "multipart/form-data") {
		return handler.updateLansiaWithImage(ctx, xMember)
	}

	// Handle regular JSON update
	var req valueobject.LansiaPayloadUpdate
	if err := ctx.BodyParser(&req); err != nil {
		logger.Error("Failed to parse request body: " + err.Error())
		return response.BadRequest(ctx, "Invalid request payload")
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

func (handler *LansiaHandler) updateLansiaWithImage(ctx *fiber.Ctx, xMember string) error {
	log.Println("Update lansia with image request received")

	// Get UUID from form data
	uuid := ctx.FormValue("uuid")
	if uuid == "" {
		logger.Error("UUID is required in form data")
		return response.BadRequest(ctx, "UUID is required")
	}

	// Parse form data for other fields
	var updateData entity.Lansia
	var hasUpdates bool

	if nama := ctx.FormValue("nama"); nama != "" {
		updateData.Nama = &nama
		hasUpdates = true
	}

	if level := ctx.FormValue("level"); level != "" {
		updateData.Level = &level
		hasUpdates = true
	}

	if caregiver := ctx.FormValue("caregiver"); caregiver != "" {
		updateData.Caregiver = &caregiver
		hasUpdates = true
	}

	// Handle image upload if present
	file, err := ctx.FormFile("image")
	if err == nil {
		// Process image upload
		imagePath, err := image.ProcessImageUpload(file, uuid)
		if err != nil {
			logger.Error("Failed to process image upload: " + err.Error())
			return response.Error(ctx, "Failed to upload image: "+err.Error())
		}
		updateData.PathGambar = &imagePath
		hasUpdates = true
		log.Printf("Image uploaded successfully: %s", imagePath)
	}

	// If no updates provided
	if !hasUpdates {
		logger.Error("No update data provided")
		return response.BadRequest(ctx, "No update data provided")
	}

	// Create update request
	updateReq := valueobject.LansiaPayloadUpdate{
		Data: valueobject.LansiaDataUpdate{
			Param: valueobject.Lansia{
				Lansia: entity.Lansia{
					UUID: uuid,
				},
			},
			Body: valueobject.Lansia{
				Lansia: updateData,
			},
		},
		User: xMember,
	}

	err = handler.lltUsecase.UpdateLansia(updateReq)
	if err != nil {
		// If database update fails and image was uploaded, delete it
		if updateData.PathGambar != nil {
			image.DeleteImage(*updateData.PathGambar)
		}
		logger.Error("Failed to update lansia: " + err.Error())
		return response.Error(ctx, "Failed to update lansia: "+err.Error())
	}

	logger.Info("Lansia updated successfully with UUID: " + uuid)
	return response.Success(ctx, map[string]any{
		"message":    "Lansia updated successfully",
		"uuid":       uuid,
		"image_path": updateData.PathGambar,
	})
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

func (handler *LansiaHandler) StoreLansiaComplete(ctx *fiber.Ctx) error {
	log.Println("Store lansia complete request received")

	userMember, err := header.ValidateAndExtractXMember(ctx)
	if err != nil {
		return response.BadRequest(ctx, err.Error())
	}

	var req valueobject.LansiaCompletePayloadInsert
	if err := ctx.BodyParser(&req); err != nil {
		logger.Error("Failed to parse request body: " + err.Error())
		return response.BadRequest(ctx, "Invalid request payload")
	}

	req.User = userMember

	results, err := handler.lltUsecase.StoreLansiaComplete(req)
	if err != nil {
		logger.Error("Failed to store lansia complete: " + err.Error())
		return response.Error(ctx, "Failed to store lansia complete: " + err.Error())
	}

	logger.Info("Lansia complete stored successfully")
	return response.Success(ctx, valueobject.LansiaListResponse{
		Message: "Lansia created successfully with QR code",
		Data:    results,
		Total:   len(results),
	})
}

func (handler *LansiaHandler) UploadLansiaImage(ctx *fiber.Ctx) error {
	log.Println("Upload lansia image request received")

	// Validate X-Member header
	xMember, err := header.ValidateAndExtractXMember(ctx)
	if err != nil {
		logger.Error("X-Member header validation failed: " + err.Error())
		return err
	}

	// Get UUID from query parameter
	uuid := ctx.Query("uuid")
	if uuid == "" {
		logger.Error("UUID query parameter is required")
		return response.BadRequest(ctx, "UUID query parameter is required")
	}

	// Get uploaded file
	file, err := ctx.FormFile("image")
	if err != nil {
		logger.Error("Failed to get uploaded file: " + err.Error())
		return response.BadRequest(ctx, "Image file is required")
	}

	// Process image upload with compression
	imagePath, err := image.ProcessImageUpload(file, uuid)
	if err != nil {
		logger.Error("Failed to process image upload: " + err.Error())
		return response.Error(ctx, "Failed to upload image: " + err.Error())
	}

	// Update lansia record with image path
	updateReq := valueobject.LansiaPayloadUpdate{
		Data: valueobject.LansiaDataUpdate{
			Param: valueobject.Lansia{
				Lansia: entity.Lansia{
					UUID: uuid,
				},
			},
			Body: valueobject.Lansia{
				Lansia: entity.Lansia{
					PathGambar: &imagePath,
				},
			},
		},
		User: xMember,
	}

	err = handler.lltUsecase.UpdateLansia(updateReq)
	if err != nil {
		// If database update fails, delete uploaded image
		image.DeleteImage(imagePath)
		logger.Error("Failed to update lansia with image path: " + err.Error())
		return response.Error(ctx, "Failed to update lansia with image")
	}

	logger.Info("Lansia image uploaded successfully for UUID: " + uuid)
	return response.Success(ctx, map[string]interface{}{
		"message":    "Image uploaded successfully",
		"uuid":       uuid,
		"image_path": imagePath,
	})
}
