package masterdata

import (
	"svc-llt-golang/valueobject"
)

type Usecase interface {
	// User usecases
	Login(username, password string) (string, error)
	GetAllUser(param map[string]interface{}) ([]valueobject.User, error)
	GetOneUser(param map[string]interface{}) (valueobject.User, error)
	CreateUser(user valueobject.User) error
	UpdateUser(param map[string]interface{}, data map[string]interface{}) error
	DeleteUser(param map[string]interface{}) error
}