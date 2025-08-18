package valueobject

import "svc-llt-golang/entity"

type Identitas struct {
	entity.Identitas
	entity.StandardKey
	entity.Pagination
	entity.Time
}

func IdentitasFromEntity(identitasEntity entity.Identitas) Identitas {
	return Identitas{
		Identitas: identitasEntity,
	}
}

type IdentitasPayloadInsert struct {
	Data []Identitas `json:"data" binding:"required"`
	User string
}

type IdentitasPayloadUpdate struct {
	Data IdentitasDataUpdate `json:"data" binding:"required"`
	User string
}

type IdentitasDataUpdate struct {
	Param Identitas `json:"param" binding:"required"`
	Body  Identitas `json:"body" binding:"required"`
}

type IdentitasPayloadDelete struct {
	Param []Identitas `json:"param" binding:"required"`
	User  string
}

type IdentitasListResponse struct {
	Message string      `json:"message"`
	Data    []Identitas `json:"data"`
	Total   int         `json:"total"`
}