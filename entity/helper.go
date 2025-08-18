package entity

import "time"

type StandardKey struct {
	ID uint64 `json:"id"`
}

type Pagination struct {
	Count int `json:"-"`
}

type Time struct {
	UserInput     string     `json:"user_input,omitempty" form:"user_input"`
	TanggalInput  *time.Time `json:"tgl_input,omitempty"`
	UserUpdate    string     `json:"user_update,omitempty"`
	TanggalUpdate *time.Time `json:"tgl_update,omitempty"`
}