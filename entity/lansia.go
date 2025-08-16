package entity

import "time"

type Lansia struct {
	ID           int64   `gorm:"primaryKey;autoIncrement:false"`
	UUID         string  `gorm:"type:varchar(255);uniqueIndex;not null"`
	Nama         *string `gorm:"column:nama;type:varchar(255)"`
	Age          int     `gorm:"not null"`
	Status       string  `gorm:"type:varchar(100);not null"`
	IdIdentitas  int64   `gorm:"column:id_identitas"`
	IdAlamat     int64   `gorm:"column:id_alamat"`
	FlagAktif    bool    `gorm:"column:flag_aktif;default:true"`
	FlagDelete   bool    `gorm:"column:flag_delete;default:false"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	UserInput    string  `gorm:"column:user_input;type:varchar(255);not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserUpdate   *string `gorm:"column:user_update;type:varchar(255);constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Lansia) TableName() string {
	return "lansia"
}
