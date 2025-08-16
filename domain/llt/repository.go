package llt

import (
	"svc-llt-golang/valueobject"
)

type Repository interface {
	// Lansia repositories
	GetAllLansia(param map[string]interface{}) ([]valueobject.Lansia, error)
	GetOneLansia(param map[string]interface{}) (valueobject.Lansia, error)
	CreateLansia(uuid, name string, age int, status string) error
	UpdateLansia(param map[string]interface{}, data map[string]interface{}) error
	DeleteLansia(param map[string]interface{}) error
}