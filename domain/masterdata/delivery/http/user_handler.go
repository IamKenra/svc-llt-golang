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
	userUsecase masterdata.Usecase
}

func NewUserHandler(userUC masterdata.Usecase) *UserHandler {
	return &UserHandler{userUsecase: userUC}
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	log.Println("Login request received")
	var req valueobject.UserLoginRequest

	if err := c.BodyParser(&req); err != nil {
		logger.Error("Failed to parse request body: " + err.Error())
		return response.BadRequest(c, "Invalid request payload")
	}

	token, err := h.userUsecase.Login(req.Username, req.Password)
	if err != nil {
		logger.Error("Invalid credentials for user: " + req.Username + " Wrong password atau username")
		return response.Error(c, "Invalid credentials Wrong password atau username")
	}

	logger.Info("User logged in successfully: " + req.Username)
	return response.Success(c, valueobject.UserLoginResponse{
		Message: "Login successful",
		Token:   token,
	})
}