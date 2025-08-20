package auth

import (
	"svc-llt-golang/utils/response"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Example of manual JWT validation in a handler
func ExampleHandlerWithManualAuth(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract token from Authorization header
		authHeader := c.Get("Authorization")
		token, err := ExtractTokenFromHeader(authHeader)
		if err != nil {
			return response.Unauthorized(c, err.Error())
		}

		// Validate token against database
		authUser, err := ValidateToken(token, db)
		if err != nil {
			return response.Unauthorized(c, err.Error())
		}

		// Now you have authenticated user information
		// You can use authUser.ID and authUser.Username
		
		return c.JSON(fiber.Map{
			"message": "Authentication successful",
			"user": fiber.Map{
				"id":       authUser.ID,
				"username": authUser.Username,
			},
		})
	}
}

// Example of getting user info from token without database verification
func ExampleGetUserFromTokenHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		token, err := ExtractTokenFromHeader(authHeader)
		if err != nil {
			return response.Unauthorized(c, err.Error())
		}

		// Get user info from token (no database check)
		userClaims, err := GetUserFromToken(token)
		if err != nil {
			return response.Unauthorized(c, err.Error())
		}

		return c.JSON(fiber.Map{
			"message":  "Token parsed successfully",
			"username": userClaims.Username,
			"claims":   userClaims.MapClaims,
		})
	}
}