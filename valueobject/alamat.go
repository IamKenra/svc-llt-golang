package valueobject

import "svc-llt-golang/entity"

type Alamat struct {
	entity.Alamat
	entity.StandardKey
	entity.Pagination
	entity.Time
}

func AlamatFromEntity(alamatEntity entity.Alamat) Alamat {
	return Alamat{
		Alamat: alamatEntity,
	}
}

type AlamatPayloadInsert struct {
	Data []Alamat `json:"data" binding:"required"`
	User string
}

type AlamatPayloadUpdate struct {
	Data AlamatDataUpdate `json:"data" binding:"required"`
	User string
}

type AlamatDataUpdate struct {
	Param Alamat `json:"param" binding:"required"`
	Body  Alamat `json:"body" binding:"required"`
}

type AlamatPayloadDelete struct {
	Param []Alamat `json:"param" binding:"required"`
	User  string
}

type AlamatListResponse struct {
	Message string   `json:"message"`
	Data    []Alamat `json:"data"`
	Total   int      `json:"total"`
}