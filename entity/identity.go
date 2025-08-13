package entity

import "time"

type Identity struct {
	ID        uint      `gorm:"primaryKey"`
	UUID      string    `gorm:"uniqueIndex;not null"`
	UserID    uint      `gorm:"not null"`
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