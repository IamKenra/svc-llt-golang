package http

import (
	"log"
	"svc-llt-golang/domain/masterdata"
	"svc-llt-golang/utils/logger"
	"svc-llt-golang/utils/response"
	"svc-llt-golang/valueobject"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	masterdataUsecase masterdata.Usecase
}

func NewUserHandler(masterdataUC masterdata.Usecase) *UserHandler {
	return &UserHandler{masterdataUsecase: masterdataUC}
}

func (handler *UserHandler) Login(ctx *fiber.Ctx) error {
	log.Println("Login request received")
	var req valueobject.UserLoginRequest

	if err := ctx.BodyParser(&req); err != nil {
		logger.Error("Failed to parse request body: " + err.Error())
		return response.BadRequest(ctx, "Invalid request payload")
	}

	token, err := handler.masterdataUsecase.Login(req.Username, req.Password)
	if err != nil {
		logger.Error("Invalid credentials for user: " + req.Username + " Wrong password atau username")
		return response.Error(ctx, "Invalid credentials Wrong password atau username")
	}

	logger.Info("User logged in successfully: " + req.Username)
	return response.Success(ctx, valueobject.UserLoginResponse{
		Message: "Login successful",
		Token:   token,
	})
}

func (handler *UserHandler) Register(ctx *fiber.Ctx) error {
	log.Println("Register request received")
	var req valueobject.UserRegisterRequest

	if err := ctx.BodyParser(&req); err != nil {
		logger.Error("Failed to parse request body: " + err.Error())
		return response.BadRequest(ctx, "Invalid request payload")
	}

	userUUID, err := handler.masterdataUsecase.Register(req)
	if err != nil {
		logger.Error("Registration failed: " + err.Error())
		return response.Error(ctx, err.Error())
	}

	logger.Info("User registered successfully: " + req.Username)
	return response.Success(ctx, valueobject.UserRegisterResponse{
		Message: "User registered successfully",
		UUID:    userUUID,
	})
}
