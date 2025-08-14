package main

import (
	"log"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"svc-llt-golang/entity"
)

func main() {
	_ = godotenv.Load()

	dsn := "admin:bismillah@1@tcp(localhost:3306)/it_asset?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
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
	if err := db.Model(&entity.User{}).Where("username = ?", username).
		Update("password", string(hashedPassword)).Error; err != nil {
		log.Fatal("Failed to update password:", err)
	}

	log.Printf("Password untuk user '%s' berhasil diupdate.\n", username)
}
