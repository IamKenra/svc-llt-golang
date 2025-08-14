package usecase

import (
	"svc-llt-golang/valueobject"
)

func (llt lltUsecase) GetAllElderlyCare(param map[string]interface{}) ([]valueobject.ElderlyCare, error) {
	return llt.repository.GetAllElderlyCare(param)
}

func (llt lltUsecase) GetOneElderlyCare(param map[string]interface{}) (valueobject.ElderlyCare, error) {
	return llt.repository.GetOneElderlyCare(param)
}

func (llt lltUsecase) StoreElderlyCare(payload valueobject.ElderlyCarePayloadInsert) (valueobject.ElderlyCarePayloadInsert, error) {
	// TODO: Generate UUID and other business logic before storing
	
	err := llt.repository.CreateElderlyCare(payload.Data[0])
	return payload, err
}

func (llt lltUsecase) UpdateElderlyCare(payload valueobject.ElderlyCarePayloadUpdate) error {
	param := map[string]interface{}{
		"uuid": payload.Data.Param.UUID,
	}
	data := map[string]interface{}{
		"name":   payload.Data.Body.Name,
		"age":    payload.Data.Body.Age,
		"status": payload.Data.Body.Status,
	}
	return llt.repository.UpdateElderlyCare(param, data)
}

func (llt lltUsecase) DeleteElderlyCare(payload valueobject.ElderlyCarePayloadDelete) error {
	for _, elderlyCare := range payload.Param {
		param := map[string]interface{}{
			"uuid": elderlyCare.UUID,
		}
		err := llt.repository.DeleteElderlyCare(param)
		if err != nil {
			return err
		}
	}
	return nil
}