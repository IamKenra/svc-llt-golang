package usecase

import (
	"svc-llt-golang/valueobject"
)

func (llt lltUsecase) ProcessStoreLansia(payload valueobject.LansiaPayloadInsert) error {
	for _, x := range payload.Data {
		var nama, pathQr, level, caregiver string
		
		if x.Lansia.Nama != nil {
			nama = *x.Lansia.Nama
		}
		if x.Lansia.PathQr != nil {
			pathQr = *x.Lansia.PathQr
		}
		if x.Lansia.Level != nil {
			level = *x.Lansia.Level
		}
		if x.Lansia.Caregiver != nil {
			caregiver = *x.Lansia.Caregiver
		}
		
		err := llt.repository.CreateLansia(x.Lansia.UUID, nama, x.Lansia.IdIdentitas, x.Lansia.IdAlamat, pathQr, level, caregiver, payload.User)
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
