package masterdata

import (
	"svc-llt-golang/valueobject"
)

type Usecase interface {
	// Authentication usecases
	Login(username, password string) (string, error)
	Register(payload valueobject.AuthRegisterRequest) (string, error)

	// User usecases
	GetAllUser(param map[string]interface{}) ([]valueobject.User, error)
	GetOneUser(param map[string]interface{}) (valueobject.User, error)
	StoreUser(payload valueobject.UserPayloadInsert) (valueobject.UserPayloadInsert, error)
	UpdateUser(payload valueobject.UserPayloadUpdate) error
	DeleteUser(payload valueobject.UserPayloadDelete) error

}
