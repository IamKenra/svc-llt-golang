package usecase

import (
	"svc-llt-golang/domain/masterdata"
)

/*
*
the struct of usecases
*/
type masterdataUsecase struct {
	repository masterdata.Repository
	jwtKey     string
}

/*
*
the initiator function for usecase
*/
func NewMasterdataUsecase(repository masterdata.Repository, jwtKey string) masterdata.Usecase {
	return &masterdataUsecase{
		repository: repository,
		jwtKey:     jwtKey,
	}
}