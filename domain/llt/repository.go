package llt

import (
	"svc-llt-golang/valueobject"
)

type Repository interface {
	// Lansia repositories
	GetAllLansia(param map[string]interface{}) ([]valueobject.Lansia, error)
	GetOneLansia(param map[string]interface{}) (valueobject.Lansia, error)
	CreateLansia(uuid, nama string, idIdentitas, idAlamat int64, userInput string) error
	UpdateLansia(param map[string]interface{}, data map[string]interface{}) error
	DeleteLansia(param map[string]interface{}) error

	// Identitas repositories
	GetAllIdentitas(param map[string]interface{}) ([]valueobject.Identitas, error)
	GetOneIdentitas(param map[string]interface{}) (valueobject.Identitas, error)
	CreateIdentitas(uuid, nik, namaDepan string, namaBelakang *string, userInput string) error
	UpdateIdentitas(param map[string]interface{}, data map[string]interface{}) error
	DeleteIdentitas(param map[string]interface{}) error

	// Alamat repositories
	GetAllAlamat(param map[string]interface{}) ([]valueobject.Alamat, error)
	GetOneAlamat(param map[string]interface{}) (valueobject.Alamat, error)
	CreateAlamat(uuid string, alamat *string, userInput string) error
	UpdateAlamat(param map[string]interface{}, data map[string]interface{}) error
	DeleteAlamat(param map[string]interface{}) error
}