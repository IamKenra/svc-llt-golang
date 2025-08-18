package usecase

import (
	"time"
	"svc-llt-golang/entity"
	"svc-llt-golang/utils/qrcode"
	"svc-llt-golang/valueobject"

	"github.com/google/uuid"
)

func (llt lltUsecase) GetAllLansia(param map[string]interface{}) ([]valueobject.Lansia, error) {
	return llt.repository.GetAllLansia(param)
}

func (llt lltUsecase) GetOneLansia(param map[string]interface{}) (valueobject.Lansia, error) {
	return llt.repository.GetOneLansia(param)
}

// Plain store - requires existing identitas & alamat IDs
func (llt lltUsecase) StoreLansia(payload valueobject.LansiaPayloadInsert) (valueobject.LansiaPayloadInsert, error) {
	for i := range payload.Data {
		// Only generate UUID if not already provided (for orchestrated creation)
		if payload.Data[i].Lansia.UUID == "" {
			payload.Data[i].Lansia.UUID = uuid.New().String()
		}
	}

	err := llt.ProcessStoreLansia(payload)
	if err != nil {
		return payload, err
	}

	// Retrieve inserted data to get correct IDs
	for i := range payload.Data {
		param := map[string]interface{}{
			"uuid": payload.Data[i].Lansia.UUID,
		}
		insertedData, err := llt.repository.GetOneLansia(param)
		if err == nil {
			payload.Data[i] = insertedData
		}
	}

	return payload, err
}

// Orchestrated store - creates identitas, alamat, then lansia with QR code
func (llt lltUsecase) StoreLansiaComplete(payload valueobject.LansiaCompletePayloadInsert) ([]valueobject.Lansia, error) {
	var results []valueobject.Lansia
	
	for _, data := range payload.Data {
		// Step 1: Create Identitas using existing function
		identitasPayload := valueobject.IdentitasPayloadInsert{
			Data: []valueobject.Identitas{
				{
					Identitas: entity.Identitas{
						Nik:              data.Identitas.Nik,
						NamaDepan:        data.Identitas.NamaDepan,
						NamaBelakang:     data.Identitas.NamaBelakang,
						Agama:            data.Identitas.Agama,
						StatusPerkawinan: data.Identitas.StatusPerkawinan,
						Pekerjaan:        data.Identitas.Pekerjaan,
					},
				},
			},
			User: payload.User,
		}
		
		// Parse TglLahir if provided
		if data.Identitas.TglLahir != nil && *data.Identitas.TglLahir != "" {
			if parsed, err := time.Parse("2006-01-02", *data.Identitas.TglLahir); err == nil {
				identitasPayload.Data[0].Identitas.TglLahir = &parsed
			}
		}
		
		identitasResult, err := llt.StoreIdentitas(identitasPayload)
		if err != nil {
			return nil, err
		}
		identitasID := identitasResult.Data[0].Identitas.ID
		
		// Step 2: Create Alamat using existing function
		alamatPayload := valueobject.AlamatPayloadInsert{
			Data: []valueobject.Alamat{
				{
					Alamat: entity.Alamat{
						Alamat:    data.Alamat.Alamat,
						Rt:        data.Alamat.Rt,
						Rw:        data.Alamat.Rw,
						Kelurahan: data.Alamat.Kelurahan,
						Kecamatan: data.Alamat.Kecamatan,
						Kota:      data.Alamat.Kota,
						Kodepos:   data.Alamat.Kodepos,
					},
				},
			},
			User: payload.User,
		}
		
		alamatResult, err := llt.StoreAlamat(alamatPayload)
		if err != nil {
			return nil, err
		}
		alamatID := alamatResult.Data[0].Alamat.ID
		
		// Step 3: Generate UUID and QR Code
		lansiaUUID := uuid.New().String()
		qrPath, err := llt.GenerateQRCode(lansiaUUID)
		if err != nil {
			return nil, err
		}
		
		// Step 4: Create Lansia using existing function
		lansiaPayload := valueobject.LansiaPayloadInsert{
			Data: []valueobject.Lansia{
				{
					Lansia: entity.Lansia{
						UUID:        lansiaUUID,
						IdIdentitas: identitasID,
						IdAlamat:    alamatID,
						Nama:        data.Lansia.Nama,
						PathQr:      &qrPath,
						Level:       data.Lansia.Level,
						Caregiver:   data.Lansia.Caregiver,
					},
				},
			},
			User: payload.User,
		}
		
		lansiaResult, err := llt.StoreLansia(lansiaPayload)
		if err != nil {
			return nil, err
		}
		
		// Convert result to response format
		for _, lansia := range lansiaResult.Data {
			results = append(results, lansia)
		}
	}
	
	return results, nil
}

func (llt lltUsecase) UpdateLansia(payload valueobject.LansiaPayloadUpdate) error {
	return llt.ProcessUpdateLansia(payload)
}

func (llt lltUsecase) DeleteLansia(payload valueobject.LansiaPayloadDelete) error {
	return llt.ProcessDeleteLansia(payload)
}

// GenerateQRCode generates QR code from UUID and returns file path
func (llt lltUsecase) GenerateQRCode(uuid string) (string, error) {
	return qrcode.GenerateQRCode(uuid)
}
