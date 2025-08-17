package usecase

import (
	"svc-llt-golang/valueobject"

	"github.com/google/uuid"
)

func (llt lltUsecase) GetAllIdentitas(param map[string]interface{}) ([]valueobject.Identitas, error) {
	return llt.repository.GetAllIdentitas(param)
}

func (llt lltUsecase) GetOneIdentitas(param map[string]interface{}) (valueobject.Identitas, error) {
	return llt.repository.GetOneIdentitas(param)
}

func (llt lltUsecase) StoreIdentitas(payload valueobject.IdentitasPayloadInsert) (valueobject.IdentitasPayloadInsert, error) {
	for i := range payload.Data {
		payload.Data[i].Identitas.UUID = uuid.New().String()
	}

	err := llt.ProcessStoreIdentitas(payload)
	return payload, err
}

func (llt lltUsecase) UpdateIdentitas(payload valueobject.IdentitasPayloadUpdate) error {
	return llt.ProcessUpdateIdentitas(payload)
}

func (llt lltUsecase) DeleteIdentitas(payload valueobject.IdentitasPayloadDelete) error {
	return llt.ProcessDeleteIdentitas(payload)
}