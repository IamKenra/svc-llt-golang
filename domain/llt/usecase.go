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
	UpdateLansia(payload valueobject.LansiaPayloadUpdate) error
	DeleteLansia(payload valueobject.LansiaPayloadDelete) error
}