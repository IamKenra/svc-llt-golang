package llt

import (
	"svc-llt-golang/valueobject"
)

/*
*
why there's only one usecase interface while there can more than one repository interface?...
... because, at DDD (Domain Design Driven), there's only one set of usecase and...
... the function name inside the usecase should be unique and represent the business process...
... tl;dr: function name is telling what exactly are they doing.
*/
type Usecase interface {
	// Lansia usecases
	GetAllLansia(param map[string]interface{}) ([]valueobject.Lansia, error)
	GetOneLansia(param map[string]interface{}) (valueobject.Lansia, error)
	StoreLansia(payload valueobject.LansiaPayloadInsert) (valueobject.LansiaPayloadInsert, error)
	StoreLansiaComplete(payload valueobject.LansiaCompletePayloadInsert) ([]valueobject.Lansia, error)
	UpdateLansia(payload valueobject.LansiaPayloadUpdate) error
	DeleteLansia(payload valueobject.LansiaPayloadDelete) error

	// Identitas usecases
	GetAllIdentitas(param map[string]interface{}) ([]valueobject.Identitas, error)
	GetOneIdentitas(param map[string]interface{}) (valueobject.Identitas, error)
	StoreIdentitas(payload valueobject.IdentitasPayloadInsert) (valueobject.IdentitasPayloadInsert, error)
	UpdateIdentitas(payload valueobject.IdentitasPayloadUpdate) error
	DeleteIdentitas(payload valueobject.IdentitasPayloadDelete) error

	// Alamat usecases
	GetAllAlamat(param map[string]interface{}) ([]valueobject.Alamat, error)
	GetOneAlamat(param map[string]interface{}) (valueobject.Alamat, error)
	StoreAlamat(payload valueobject.AlamatPayloadInsert) (valueobject.AlamatPayloadInsert, error)
	UpdateAlamat(payload valueobject.AlamatPayloadUpdate) error
	DeleteAlamat(payload valueobject.AlamatPayloadDelete) error
}