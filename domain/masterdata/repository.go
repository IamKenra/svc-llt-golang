package masterdata

import (
	"svc-llt-golang/valueobject"
)

type Repository interface {
	// Health check repositories
	HealthCheck() error

	// User repositories
	GetAllUser(param map[string]interface{}) ([]valueobject.User, error)
	GetOneUser(param map[string]interface{}) (valueobject.User, error)
	FindByUsername(username string) (valueobject.User, error)
	FindByUUID(uuid string) (valueobject.User, error)
	CreateUser(user valueobject.User) error
	UpdateUser(param map[string]interface{}, data map[string]interface{}) error
	DeleteUser(param map[string]interface{}) error
}