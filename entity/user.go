package entity

import (
	"time"
)

type Auth struct {
	ID         int64     `gorm:"primaryKey;autoIncrement:false"`
	Username   string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password   string    `gorm:"type:varchar(255);not null"`
	TglInput   time.Time `gorm:"column:tgl_input"`
	TglUpdate  time.Time `gorm:"column:tgl_update"`
	UserInput  string    `gorm:"column:user_input"`
	UserDelete *string   `gorm:"column:user_delete"`
}

func (Auth) TableName() string {
	return "auth"
}

type User struct {
	ID         int64     `gorm:"primaryKey;autoIncrement:false"`
	UUID       string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	Nama       string    `gorm:"type:varchar(255);not null"`
	Email      string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	TglInput   time.Time `gorm:"column:tgl_input"`
	TglUpdate  time.Time `gorm:"column:tgl_update"`
	UserInput  string    `gorm:"column:user_input"`
	UserUpdate *string   `gorm:"column:user_update"`
}

func (User) TableName() string {
	return "user"
}
