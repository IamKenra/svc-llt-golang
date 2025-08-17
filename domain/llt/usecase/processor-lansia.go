package usecase

import (
	"svc-llt-golang/valueobject"
)

func (llt lltUsecase) ProcessStoreLansia(payload valueobject.LansiaPayloadInsert) error {
	for _, x := range payload.Data {
		var nama string
		if x.Lansia.Nama != nil {
			nama = *x.Lansia.Nama
		}
		err := llt.repository.CreateLansia(x.Lansia.UUID, nama, x.Lansia.IdIdentitas, x.Lansia.IdAlamat, payload.User)
		if err != nil {
			return err
		}
	}
	return nil
}

func (llt lltUsecase) ProcessUpdateLansia(payload valueobject.LansiaPayloadUpdate) error {
	param := map[string]interface{}{
		"uuid": payload.Data.Param.Lansia.UUID,
	}
	data := map[string]interface{}{}

	if payload.User != "" {
		data["user_update"] = payload.User
	}

	return llt.repository.UpdateLansia(param, data)
}

func (llt lltUsecase) ProcessDeleteLansia(payload valueobject.LansiaPayloadDelete) error {
	for _, x := range payload.Param {
		param := map[string]interface{}{
			"uuid": x.Lansia.UUID,
		}
		err := llt.repository.DeleteLansia(param)
		if err != nil {
			return err
		}
	}
	return nil
}
