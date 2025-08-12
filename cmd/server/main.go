package main

import (
	"log"
	"os"
	v1 "svc-itasset-golang/api/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"svc-itasset-golang/config"
	"svc-itasset-golang/internal/middleware"
)

func main() {
	_ = godotenv.Load()

	dsn := config.GetDatabaseDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	app := fiber.New()

	app.Use(middleware.CORSConfig())

	apiGroup := app.Group("/itasset")
	v1.RegisterRoutes(apiGroup, db, os.Getenv("JWT_SECRET"))

	log.Fatal(app.Listen(":3000"))
}
