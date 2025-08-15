package main

import (
	"log"
	"os"
	"svc-llt-golang/domain/masterdata/delivery/http"
	"svc-llt-golang/entity"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"svc-llt-golang/config"
	"svc-llt-golang/utils/middleware"
)

func main() {
	_ = godotenv.Load()

	dsn := config.GetDatabaseDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Drop existing tables and recreate them to match our entities
	err = db.Migrator().DropTable(&entity.Auth{}, &entity.User{})
	if err != nil {
		log.Printf("Warning: Failed to drop existing tables: %v", err)
	}

	// Auto-migrate tables based on entities
	err = db.AutoMigrate(&entity.Auth{}, &entity.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	app := fiber.New()

	app.Use(middleware.CORSConfig())

	// Register routes with /llt-svc prefix
	apiGroup := app.Group("/llt-svc")
	http.RegisterRoutes(apiGroup, db, os.Getenv("JWT_SECRET"))

	// Add a test route to verify routing works
	apiGroup.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "test endpoint working"})
	})

	log.Println("Routes registered. Starting server...")
	log.Fatal(app.Listen(":3000"))
}
