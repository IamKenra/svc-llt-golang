package valueobject

import "svc-llt-golang/entity"

type User struct {
	entity.User
	entity.StandardKey
	entity.Pagination
	entity.Time
}

func UserFromEntity(e entity.User) User {
	return User{
		User: e,
	}
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