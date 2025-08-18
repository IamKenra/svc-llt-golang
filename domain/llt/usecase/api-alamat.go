package usecase

import (
	"svc-llt-golang/valueobject"

	"github.com/google/uuid"
)

func (llt lltUsecase) GetAllAlamat(param map[string]interface{}) ([]valueobject.Alamat, error) {
	return llt.repository.GetAllAlamat(param)
}

func (llt lltUsecase) GetOneAlamat(param map[string]interface{}) (valueobject.Alamat, error) {
	return llt.repository.GetOneAlamat(param)
}

func (llt lltUsecase) StoreAlamat(payload valueobject.AlamatPayloadInsert) (valueobject.AlamatPayloadInsert, error) {
	for i := range payload.Data {
		payload.Data[i].Alamat.UUID = uuid.New().String()
	}

	err := llt.ProcessStoreAlamat(payload)
	if err != nil {
		return payload, err
	}

	// Retrieve inserted data to get correct IDs
	for i := range payload.Data {
		param := map[string]interface{}{
			"uuid": payload.Data[i].Alamat.UUID,
		}
		insertedData, err := llt.repository.GetOneAlamat(param)
		if err == nil {
			payload.Data[i] = insertedData
		}
	}

	return payload, err
}

func (llt lltUsecase) UpdateAlamat(payload valueobject.AlamatPayloadUpdate) error {
	return llt.ProcessUpdateAlamat(payload)
}

func (llt lltUsecase) DeleteAlamat(payload valueobject.AlamatPayloadDelete) error {
	return llt.ProcessDeleteAlamat(payload)
}