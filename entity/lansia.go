package entity

import "time"

type Lansia struct {
	ID           int64   `gorm:"primaryKey;autoIncrement:false"`
	UUID         string  `gorm:"type:varchar(255);uniqueIndex;not null"`
	IdIdentitas  int64   `gorm:"column:id_identitas;not null"`
	IdAlamat     int64   `gorm:"column:id_alamat;not null"`
	Nama         *string `gorm:"column:nama;type:varchar(255)"`
	PathGambar   *string `gorm:"column:path_gambar;type:varchar(255)"`
	PathQr       *string `gorm:"column:path_qr;type:varchar(255)"`
	Level        *string `gorm:"column:level;type:varchar(50)"`
	Caregiver    *string `gorm:"column:caregiver;type:varchar(150)"`
	FlagAktif    bool    `gorm:"column:flag_aktif;default:true"`
	FlagDelete   bool    `gorm:"column:flag_delete;default:false"`
	UserInput    *string `gorm:"column:user_input;type:varchar(50)"`
	UserUpdate   *string `gorm:"column:user_update;type:varchar(50)"`
	TglInput     time.Time `gorm:"column:tgl_input;default:CURRENT_TIMESTAMP"`
	TglUpdate    time.Time `gorm:"column:tgl_update;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (Lansia) TableName() string {
	return "lansia"
}
