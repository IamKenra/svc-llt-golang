package valueobject

import "svc-llt-golang/entity"

type Lansia struct {
	entity.Lansia
	// Add additional fields for API responses if needed
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