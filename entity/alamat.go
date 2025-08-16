package entity

import "time"

type Alamat struct {
	ID         int64     `gorm:"primaryKey;autoIncrement:false"`
	UUID       string    `gorm:"type:varchar(255);uniqueIndex:uq_alamat_uuid;not null"`
	Alamat     *string   `gorm:"type:varchar(255)"`
	Rt         *string   `gorm:"type:varchar(5)"`
	Rw         *string   `gorm:"type:varchar(5)"`
	Kelurahan  *string   `gorm:"type:varchar(100)"`
	Kecamatan  *string   `gorm:"type:varchar(100)"`
	Kota       *string   `gorm:"type:varchar(100)"`
	Kodepos    *string   `gorm:"type:varchar(10)"`
	FlagAktif  bool      `gorm:"column:flag_aktif;type:tinyint(1);not null;default:1"`
	UserInput  *string   `gorm:"column:user_input;type:varchar(50)"`
	UserUpdate *string   `gorm:"column:user_update;type:varchar(50)"`
	TglInput   time.Time `gorm:"column:tgl_input;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	TglUpdate  time.Time `gorm:"column:tgl_update;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (Alamat) TableName() string {
	return "alamat"
}