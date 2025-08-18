package valueobject

import "svc-llt-golang/entity"

type Lansia struct {
	entity.Lansia
	entity.StandardKey
	entity.Pagination
	entity.Time
	Age int `json:"age,omitempty"` // Calculated from identitas.tgl_lahir
}

func LansiaFromEntity(lansiaEntity entity.Lansia) Lansia {
	return Lansia{
		Lansia: lansiaEntity,
	}
}

// Plain payload - direct database mapping (requires existing identitas & alamat IDs)
type LansiaPayloadInsert struct {
	Data []Lansia `json:"data" binding:"required"`
	User string
}

// Orchestrated payload - complete lansia creation with auto QR generation
type LansiaCompletePayloadInsert struct {
	Data []LansiaCompleteData `json:"data" binding:"required"`
	User string
}

type LansiaCompleteData struct {
	// Identitas data (will be created first)
	Identitas IdentitasData `json:"identitas" binding:"required"`
	// Alamat data (will be created first)  
	Alamat    AlamatData    `json:"alamat" binding:"required"`
	// Lansia data (UUID & QR will be auto-generated)
	Lansia    LansiaData    `json:"lansia" binding:"required"`
}

// Simplified data structures for orchestrated creation
type IdentitasData struct {
	Nik              string  `json:"nik" binding:"required"`
	NamaDepan        string  `json:"nama_depan" binding:"required"`
	NamaBelakang     *string `json:"nama_belakang"`
	TglLahir         *string `json:"tgl_lahir"` // Format: "2006-01-02"
	Agama            *string `json:"agama"`
	StatusPerkawinan *string `json:"status_perkawinan"`
	Pekerjaan        *string `json:"pekerjaan"`
}

type AlamatData struct {
	Alamat    *string `json:"alamat"`
	Rt        *string `json:"rt"`
	Rw        *string `json:"rw"`
	Kelurahan *string `json:"kelurahan"`
	Kecamatan *string `json:"kecamatan"`
	Kota      *string `json:"kota"`
	Kodepos   *string `json:"kodepos"`
}

type LansiaData struct {
	Nama      *string `json:"nama"`
	Level     *string `json:"level"`
	Caregiver *string `json:"caregiver"`
	// NOTE: UUID, PathQr, dan PathGambar akan di-generate/handle terpisah
	//       PathGambar di-handle saat update profile
}

type LansiaPayloadUpdate struct {
	Data LansiaDataUpdate `json:"data" binding:"required"`
	User string
}

type LansiaDataUpdate struct {
	Param Lansia `json:"param" binding:"required"`
	Body  Lansia `json:"body" binding:"required"`
}

type LansiaPayloadDelete struct {
	Param []Lansia `json:"param" binding:"required"`
	User  string
}

type LansiaListResponse struct {
	Message string   `json:"message"`
	Data    []Lansia `json:"data"`
	Total   int      `json:"total"`
}