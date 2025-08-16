package usecase

import (
	"svc-llt-golang/valueobject"

	"github.com/google/uuid"
)

func (llt lltUsecase) GetAllLansia(param map[string]interface{}) ([]valueobject.Lansia, error) {
	return llt.repository.GetAllLansia(param)
}

func (llt lltUsecase) GetOneLansia(param map[string]interface{}) (valueobject.Lansia, error) {
	return llt.repository.GetOneLansia(param)
}

func (llt lltUsecase) StoreLansia(payload valueobject.LansiaPayloadInsert) (valueobject.LansiaPayloadInsert, error) {
	for i := range payload.Data {
		payload.Data[i].UUID = uuid.New().String()
	}

	err := llt.ProcessStoreLansia(payload)
	return payload, err
}

func (llt lltUsecase) UpdateLansia(payload valueobject.LansiaPayloadUpdate) error {
	return llt.ProcessUpdateLansia(payload)
}

func (llt lltUsecase) DeleteLansia(payload valueobject.LansiaPayloadDelete) error {
	return llt.ProcessDeleteLansia(payload)
}
