package masterdata

import (
	"svc-llt-golang/valueobject"
)

type Repository interface {
	// Auth repositories
	CreateAuth(username, password string) error
	FindByUsername(username string) (valueobject.User, error)
	
	// User repositories
	GetAllUser(param map[string]interface{}) ([]valueobject.User, error)
	GetOneUser(param map[string]interface{}) (valueobject.User, error)
	FindByUUID(uuid string) (valueobject.User, error)
	CreateUser(uuid, nama, email string) error
	UpdateUser(param map[string]interface{}, data map[string]interface{}) error
	DeleteUser(param map[string]interface{}) error

}