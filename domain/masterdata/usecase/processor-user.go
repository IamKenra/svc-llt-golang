package usecase

import (
	"time"
	"svc-llt-golang/valueobject"
)

func (masterdata masterdataUsecase) ProcessStoreUser(payload valueobject.UserPayloadInsert) error {
	for _, x := range payload.Data {
		err := masterdata.repository.CreateUser(x.User.UUID, x.User.Nama, x.User.Email)
		if err != nil {
			return err
		}
	}
	return nil
}

func (masterdata masterdataUsecase) ProcessUpdateUser(payload valueobject.UserPayloadUpdate) error {
	param := map[string]interface{}{
		"uuid": payload.Data.Param.User.UUID,
	}
	data := map[string]interface{}{
		"nama":       payload.Data.Body.User.Nama,
		"email":      payload.Data.Body.User.Email,
		"tgl_update": time.Now(),
		"user_update": payload.Data.Body.User.UserUpdate,
	}

	return masterdata.repository.UpdateUser(param, data)
}

func (masterdata masterdataUsecase) ProcessDeleteUser(payload valueobject.UserPayloadDelete) error {
	for _, x := range payload.Param {
		param := map[string]interface{}{
			"uuid": x.User.UUID,
		}
		err := masterdata.repository.DeleteUser(param)
		if err != nil {
			return err
		}
	}
	return nil
}