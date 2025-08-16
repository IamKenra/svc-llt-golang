package usecase

import (
	"time"
	"svc-llt-golang/valueobject"
)

func (llt lltUsecase) ProcessStoreLansia(payload valueobject.LansiaPayloadInsert) error {
	for _, x := range payload.Data {
		err := llt.repository.CreateLansia(x.UUID, x.Name, x.Age, x.Status)
		if err != nil {
			return err
		}
	}
	return nil
}

func (llt lltUsecase) ProcessUpdateLansia(payload valueobject.LansiaPayloadUpdate) error {
	param := map[string]interface{}{
		"uuid": payload.Data.Param.UUID,
	}
	data := map[string]interface{}{
		"name":       payload.Data.Body.Name,
		"age":        payload.Data.Body.Age,
		"status":     payload.Data.Body.Status,
		"updated_at": time.Now(),
	}

	return llt.repository.UpdateLansia(param, data)
}

func (llt lltUsecase) ProcessDeleteLansia(payload valueobject.LansiaPayloadDelete) error {
	for _, x := range payload.Param {
		param := map[string]interface{}{
			"uuid": x.UUID,
		}
		err := llt.repository.DeleteLansia(param)
		if err != nil {
			return err
		}
	}
	return nil
}