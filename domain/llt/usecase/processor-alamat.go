package usecase

import (
	"svc-llt-golang/valueobject"
)

func (llt lltUsecase) ProcessStoreAlamat(payload valueobject.AlamatPayloadInsert) error {
	for _, x := range payload.Data {
		err := llt.repository.CreateAlamat(x.Alamat.UUID, x.Alamat.Alamat, payload.User)
		if err != nil {
			return err
		}
	}
	return nil
}

func (llt lltUsecase) ProcessUpdateAlamat(payload valueobject.AlamatPayloadUpdate) error {
	param := map[string]interface{}{
		"uuid": payload.Data.Param.Alamat.UUID,
	}
	data := map[string]interface{}{}
	
	if payload.User != "" {
		data["user_update"] = payload.User
	}

	return llt.repository.UpdateAlamat(param, data)
}

func (llt lltUsecase) ProcessDeleteAlamat(payload valueobject.AlamatPayloadDelete) error {
	for _, x := range payload.Param {
		param := map[string]interface{}{
			"uuid": x.Alamat.UUID,
		}
		err := llt.repository.DeleteAlamat(param)
		if err != nil {
			return err
		}
	}
	return nil
}