package main

import (
	"log"
	"os"
	"svc-llt-golang/domain/masterdata/delivery/http"

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

	app := fiber.New()

	app.Use(middleware.CORSConfig())

	apiGroup := app.Group("/itasset")
	http.RegisterRoutes(apiGroup, db, os.Getenv("JWT_SECRET"))

	log.Fatal(app.Listen(":3000"))
}
