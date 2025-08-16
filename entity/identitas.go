package entity

import "time"

type Identitas struct {
	ID               int64     `gorm:"primaryKey;autoIncrement:false"`
	UUID             string    `gorm:"type:varchar(255);uniqueIndex:uq_biodata_uuid;not null"`
	Nik              string    `gorm:"type:varchar(32);uniqueIndex:uq_biodata_nik;not null"`
	NamaDepan        string    `gorm:"column:nama_depan;type:varchar(100);not null"`
	NamaBelakang     *string   `gorm:"column:nama_belakang;type:varchar(100)"`
	TglLahir         *time.Time `gorm:"column:tgl_lahir;type:date"`
	Agama            *string   `gorm:"type:varchar(50)"`
	StatusPerkawinan *string   `gorm:"column:status_perkawinan;type:varchar(50)"`
	Pekerjaan        *string   `gorm:"type:varchar(100)"`
	FlagAktif        bool      `gorm:"column:flag_aktif;type:tinyint(1);not null;default:1"`
	UserInput        *string   `gorm:"column:user_input;type:varchar(50)"`
	UserUpdate       *string   `gorm:"column:user_update;type:varchar(50)"`
	TglInput         time.Time `gorm:"column:tgl_input;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	TglUpdate        time.Time `gorm:"column:tgl_update;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (Identitas) TableName() string {
	return "identitas"
}