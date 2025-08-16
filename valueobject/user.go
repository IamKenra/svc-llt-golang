package valueobject

import "svc-llt-golang/entity"

type User struct {
	entity.User
	// Add additional fields for API responses if needed
	Password string `json:"-"` // For internal use only, not exposed in JSON
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type UserRegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nama     string `json:"nama" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type UserRegisterResponse struct {
	Message string `json:"message"`
	UUID    string `json:"uuid"`
}

type UserPayloadInsert struct {
	Data []User `json:"data" binding:"required"`
	User string
}

type UserPayloadUpdate struct {
	Data UserDataUpdate `json:"data" binding:"required"`
	User string
}

type UserDataUpdate struct {
	Param User `json:"param" binding:"required"`
	Body  User `json:"body" binding:"required"`
}

type UserPayloadDelete struct {
	Param []User `json:"param" binding:"required"`
	User  string
}