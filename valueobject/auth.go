package valueobject

import "svc-llt-golang/entity"

type Auth struct {
	entity.Auth
	entity.StandardKey
	entity.Pagination
	entity.Time
}

func AuthFromEntity(e entity.Auth) Auth {
	return Auth{
		Auth: e,
	}
}

type AuthLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthLoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type AuthRegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nama     string `json:"nama" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type AuthRegisterResponse struct {
	Message string `json:"message"`
	UUID    string `json:"uuid"`
}

type AuthPayloadInsert struct {
	Data []Auth `json:"data" binding:"required"`
	User string
}

type AuthPayloadUpdate struct {
	Data AuthDataUpdate `json:"data" binding:"required"`
	User string
}

type AuthDataUpdate struct {
	Param Auth `json:"param" binding:"required"`
	Body  Auth `json:"body" binding:"required"`
}

type AuthPayloadDelete struct {
	Param []Auth `json:"param" binding:"required"`
	User  string
}