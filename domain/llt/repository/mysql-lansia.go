package repository

import (
	"svc-llt-golang/domain/llt"
	"svc-llt-golang/entity"
	"svc-llt-golang/utils/utils"
	"svc-llt-golang/valueobject"
	"time"

	"gorm.io/gorm"
)

type mysqlLltRepository struct {
	db *gorm.DB
}

func NewMysqlLltRepository(db *gorm.DB) llt.Repository {
	return &mysqlLltRepository{db}
}

func (db *mysqlLltRepository) GetAllLansia(param map[string]interface{}) ([]valueobject.Lansia, error) {
	var lansiaList []entity.Lansia
	query := db.db

	for key, value := range param {
		query = query.Where(key+" = ?", value)
	}

	if err := query.Find(&lansiaList).Error; err != nil {
		return nil, err
	}

	var result []valueobject.Lansia
	for _, lansia := range lansiaList {
		result = append(result, valueobject.Lansia{Lansia: lansia})
	}

	return result, nil
}

func (db *mysqlLltRepository) GetOneLansia(param map[string]interface{}) (valueobject.Lansia, error) {
	var lansia entity.Lansia
	query := db.db

	for key, value := range param {
		query = query.Where(key+" = ?", value)
	}

	if err := query.First(&lansia).Error; err != nil {
		return valueobject.Lansia{}, err
	}

	return valueobject.Lansia{Lansia: lansia}, nil
}

func (db *mysqlLltRepository) CreateLansia(uuid, nama string, idIdentitas, idAlamat int64, userInput string) error {
	randomID, err := utils.GenerateRandomID()
	if err != nil {
		return err
	}

	lansia := entity.Lansia{
		ID:          randomID,
		UUID:        uuid,
		IdIdentitas: idIdentitas,
		IdAlamat:    idAlamat,
		FlagAktif:   true,
		FlagDelete:  false,
		TglInput:    time.Now(),
		TglUpdate:   time.Now(),
	}

	// Set optional fields
	if nama != "" {
		lansia.Nama = &nama
	}
	if userInput != "" {
		lansia.UserInput = &userInput
	}
	return db.db.Create(&lansia).Error
}

func (db *mysqlLltRepository) UpdateLansia(param map[string]interface{}, data map[string]interface{}) error {
	query := db.db.Model(&entity.Lansia{})

	for key, value := range param {
		query = query.Where(key+" = ?", value)
	}

	data["tgl_update"] = time.Now()

	return query.Updates(data).Error
}

func (db *mysqlLltRepository) DeleteLansia(param map[string]interface{}) error {
	query := db.db.Model(&entity.Lansia{})

	for key, value := range param {
		query = query.Where(key+" = ?", value)
	}

	return query.Delete(&entity.Lansia{}).Error
}
