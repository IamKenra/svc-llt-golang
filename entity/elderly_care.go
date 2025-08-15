package entity

import "time"

type ElderlyCare struct {
	ID        int64     `gorm:"primaryKey;autoIncrement:false"`
	UUID      string    `gorm:"uniqueIndex;not null"`
	Name      string    `gorm:"not null"`
	Age       int       `gorm:"not null"`
	Status    string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ElderlyCare) TableName() string {
	return "elderly_care"
}