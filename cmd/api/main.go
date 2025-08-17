package main

import (
	"log"
	"os"
	masterdataHttp "svc-llt-golang/domain/masterdata/delivery/http"
	lltHttp "svc-llt-golang/domain/llt/delivery/http"

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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		CreateBatchSize:                          1000,
	})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Database connection established - using manual schema management
	log.Println("Database connected. Using manual schema management (no auto-migration)")
	log.Println("Ensure database schema matches entities defined in /entity/ directory")

	app := fiber.New()

	app.Use(middleware.CORSConfig())

	// Register routes with /llt-svc prefix
	apiGroup := app.Group("/llt-svc")
	
	// Register masterdata routes (auth)
	masterdataHttp.RegisterRoutes(apiGroup, db, os.Getenv("JWT_SECRET"))
	
	// Register LLT routes (lansia)
	lltHttp.RegisterLltRoutes(apiGroup, db)

	// Add a test route to verify routing works
	apiGroup.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "test endpoint working"})
	})

	log.Println("Routes registered. Starting server...")
	log.Fatal(app.Listen(":3000"))
}