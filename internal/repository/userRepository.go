package repository

import (
	"svc-itasset-golang/internal/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(Username string) (*domain.Users, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindByUsername(username string) (*domain.Users, error) {
	var user domain.Users
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
