package main

import (
	"log"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"svc-llt-golang/internal/domain"
)

func main() {
	_ = godotenv.Load()

	dsn := "host=localhost user=admin password=bismillah@1 dbname=it_asset port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "auth.", // untuk schema auth
			SingularTable: false,
		},
	})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	username := "admin"       // Ganti sesuai username yang kamu punya
	newPassword := "admin123" // Password baru yang ingin kamu set
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)

	// Update password
	if err := db.Model(&domain.Users{}).Where("username = ?", username).
		Update("password", string(hashedPassword)).Error; err != nil {
		log.Fatal("Failed to update password:", err)
	}

	log.Printf("Password untuk user '%s' berhasil diupdate.\n", username)
}
