package entity

import "time"

type Identity struct {
	ID        int64     `gorm:"primaryKey;autoIncrement:false"`
	UUID      string    `gorm:"uniqueIndex;not null"`
	UserID    int64     `gorm:"not null"`
	Type      string    `gorm:"not null"` // email, phone, etc.
	Value     string    `gorm:"not null"`
	Verified  bool      `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User      `gorm:"foreignKey:UserID"`
}

func (Identity) TableName() string {
	return "identities"
}