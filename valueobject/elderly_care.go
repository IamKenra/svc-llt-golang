package valueobject

import "svc-llt-golang/entity"

type ElderlyCare struct {
	entity.ElderlyCare
	// Add additional fields for API responses if needed
}

type ElderlyCarePayloadInsert struct {
	Data []ElderlyCare `json:"data" binding:"required"`
	User string
}

type ElderlyCarePayloadUpdate struct {
	Data ElderlyCareDataUpdate `json:"data" binding:"required"`
	User string
}

type ElderlyCareDataUpdate struct {
	Param ElderlyCare `json:"param" binding:"required"`
	Body  ElderlyCare `json:"body" binding:"required"`
}

type ElderlyCarePayloadDelete struct {
	Param []ElderlyCare `json:"param" binding:"required"`
}