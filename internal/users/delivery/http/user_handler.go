package http

import (
	"log"
	"svc-llt-golang/internal/users/usecase"
	"svc-llt-golang/pkg/logger"
	"svc-llt-golang/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUC usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUC}
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	log.Println("Login request received")
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

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
	return response.Success(c, map[string]string{
		"message": "Login successful",
		"token":   token,
	})
}