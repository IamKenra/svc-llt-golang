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
	// ElderlyCare usecases
	GetAllElderlyCare(param map[string]interface{}) ([]valueobject.ElderlyCare, error)
	GetOneElderlyCare(param map[string]interface{}) (valueobject.ElderlyCare, error)
	StoreElderlyCare(payload valueobject.ElderlyCarePayloadInsert) (valueobject.ElderlyCarePayloadInsert, error)
	UpdateElderlyCare(payload valueobject.ElderlyCarePayloadUpdate) error
	DeleteElderlyCare(payload valueobject.ElderlyCarePayloadDelete) error
}