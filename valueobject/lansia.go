package valueobject

import "svc-llt-golang/entity"

type Lansia struct {
	entity.Lansia
	entity.StandardKey
	entity.Pagination
	entity.Time
	Age int `json:"age,omitempty"` // Calculated from identitas.tgl_lahir
}

func LansiaFromEntity(e entity.Lansia) Lansia {
	return Lansia{
		Lansia: e,
	}
}

type LansiaPayloadInsert struct {
	Data []Lansia `json:"data" binding:"required"`
	User string
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