package llt

import (
	"svc-llt-golang/valueobject"
)

type Repository interface {
	// ElderlyCare repositories
	GetAllElderlyCare(param map[string]interface{}) ([]valueobject.ElderlyCare, error)
	GetOneElderlyCare(param map[string]interface{}) (valueobject.ElderlyCare, error)
	CreateElderlyCare(elderlyCare valueobject.ElderlyCare) error
	UpdateElderlyCare(param map[string]interface{}, data map[string]interface{}) error
	DeleteElderlyCare(param map[string]interface{}) error
}