package usecase

import (
	"svc-llt-golang/domain/llt"
)

/*
*
the struct of usecases
*/
type lltUsecase struct {
	repository llt.Repository
}

/*
*
the initiator function for usecase
*/
func NewLltUsecase(repository llt.Repository) llt.Usecase {
	return &lltUsecase{
		repository: repository,
	}
}