package usecase

import (
	"svc-llt-golang/valueobject"
)

func (llt lltUsecase) ProcessStoreIdentitas(payload valueobject.IdentitasPayloadInsert) error {
	for _, x := range payload.Data {
		err := llt.repository.CreateIdentitas(x.Identitas.UUID, x.Identitas.Nik, x.Identitas.NamaDepan, x.Identitas.NamaBelakang, payload.User)
		if err != nil {
			return err
		}
	}
	return nil
}

func (llt lltUsecase) ProcessUpdateIdentitas(payload valueobject.IdentitasPayloadUpdate) error {
	param := map[string]interface{}{
		"uuid": payload.Data.Param.Identitas.UUID,
	}
	data := map[string]interface{}{}
	
	if payload.User != "" {
		data["user_update"] = payload.User
	}

	return llt.repository.UpdateIdentitas(param, data)
}

func (llt lltUsecase) ProcessDeleteIdentitas(payload valueobject.IdentitasPayloadDelete) error {
	for _, x := range payload.Param {
		param := map[string]interface{}{
			"uuid": x.Identitas.UUID,
		}
		err := llt.repository.DeleteIdentitas(param)
		if err != nil {
			return err
		}
	}
	return nil
}