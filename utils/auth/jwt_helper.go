package auth

import (
	"errors"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type TokenClaims struct {
	Username string `json:"username"`
	jwt.MapClaims
}

type AuthUser struct {
	ID       uint64 `gorm:"column:id"`
	Username string `gorm:"column:username"`
}

// ValidateToken validates JWT token and returns user information
func ValidateToken(tokenString string, db *gorm.DB) (*AuthUser, error) {
	// Remove Bearer prefix if present
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, errors.New("JWT secret not configured")
	}

	// Parse and validate token
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return nil, errors.New("username not found in token")
	}

	// Verify user exists in database
	var authUser AuthUser
	if err := db.Table("auth").Where("username = ?", username).First(&authUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user not found in database")
		}
		return nil, errors.New("database error during token verification")
	}

	return &authUser, nil
}

// ExtractTokenFromHeader extracts token from Authorization header
func ExtractTokenFromHeader(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errors.New("authorization header is missing")
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", errors.New("authorization header format is invalid")
	}

	return strings.TrimPrefix(authHeader, "Bearer "), nil
}

// GetUserFromToken extracts user information from a valid token string
func GetUserFromToken(tokenString string) (*TokenClaims, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, errors.New("JWT secret not configured")
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return nil, errors.New("username not found in token")
	}

	return &TokenClaims{
		Username:  username,
		MapClaims: claims,
	}, nil
}