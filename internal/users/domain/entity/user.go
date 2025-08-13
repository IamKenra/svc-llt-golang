package entity

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	UUID      string `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "auth.users"
}