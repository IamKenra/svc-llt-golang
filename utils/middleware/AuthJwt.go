package middleware

import (
	"os"
	"strings"
	"svc-llt-golang/utils/logger"
	"svc-llt-golang/utils/response"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func JWTProtected(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			logger.Error("Authorization header is missing")
			return response.Unauthorized(c, "Authorization header is missing")
		}
		if !strings.HasPrefix(authHeader, "Bearer ") {
			logger.Error("Authorization header format is invalid")
			return response.Unauthorized(c, "Authorization header format is invalid")
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		secret := os.Getenv("JWT_SECRET")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired JWT",
			})
		}

		// Verify user exists in database
		claims := token.Claims.(jwt.MapClaims)
		username, ok := claims["username"].(string)
		if !ok {
			return response.Unauthorized(c, "Invalid token claims")
		}

		// Check if user exists in database
		var authUser struct {
			ID       uint64 `gorm:"column:id"`
			Username string `gorm:"column:username"`
		}
		
		if err := db.Table("auth").Where("username = ?", username).First(&authUser).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return response.Unauthorized(c, "User not found in database")
			}
			logger.Error("Database error during token verification: " + err.Error())
			return response.Unauthorized(c, "Authentication failed")
		}

		// Set user info in context
		c.Locals("username", username)
		c.Locals("user_id", authUser.ID)

		return c.Next()
	}
}
