package middleware

import (
	"os"
	"svc-llt-golang/utils/response"

	"github.com/gofiber/fiber/v2"
)

func InternalApiKeyAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		apiKey := c.Get("x-api-key")
		expectedKey := os.Getenv("INTERNAL_API_KEY")
		if apiKey == "" {
			response.BadRequest(c, "API key is required")
			return nil
		}

		if apiKey != expectedKey {
			response.Unauthorized(c, "Invalid API key")
			return nil
		}

		return c.Next()
	}
}
