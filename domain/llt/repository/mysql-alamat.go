package repository

import (
	"svc-llt-golang/entity"
	"svc-llt-golang/utils/utils"
	"svc-llt-golang/valueobject"
	"time"
)

func (db *mysqlLltRepository) GetAllAlamat(param map[string]interface{}) ([]valueobject.Alamat, error) {
	var alamatList []entity.Alamat
	query := db.db

	for key, value := range param {
		query = query.Where(key+" = ?", value)
	}

	if err := query.Find(&alamatList).Error; err != nil {
		return nil, err
	}

	var result []valueobject.Alamat
	for _, alamat := range alamatList {
		result = append(result, valueobject.AlamatFromEntity(alamat))
	}

	return result, nil
}

func (db *mysqlLltRepository) GetOneAlamat(param map[string]interface{}) (valueobject.Alamat, error) {
	var alamat entity.Alamat
	query := db.db

	for key, value := range param {
		query = query.Where(key+" = ?", value)
	}

	if err := query.First(&alamat).Error; err != nil {
		return valueobject.Alamat{}, err
	}

	return valueobject.AlamatFromEntity(alamat), nil
}

func (db *mysqlLltRepository) CreateAlamat(uuid string, alamat *string, userInput string) error {
	randomID, err := utils.GenerateRandomID()
	if err != nil {
		return err
	}

	alamatEntity := entity.Alamat{
		ID:        randomID,
		UUID:      uuid,
		FlagAktif: true,
		TglInput:  time.Now(),
		TglUpdate: time.Now(),
	}

	if alamat != nil {
		alamatEntity.Alamat = alamat
	}
	if userInput != "" {
		alamatEntity.UserInput = &userInput
	}

	return db.db.Create(&alamatEntity).Error
}

func (db *mysqlLltRepository) UpdateAlamat(param map[string]interface{}, data map[string]interface{}) error {
	query := db.db.Model(&entity.Alamat{})

	for key, value := range param {
		query = query.Where(key+" = ?", value)
	}

	data["tgl_update"] = time.Now()
	return query.Updates(data).Error
}

func (db *mysqlLltRepository) DeleteAlamat(param map[string]interface{}) error {
	query := db.db.Model(&entity.Alamat{})

	for key, value := range param {
		query = query.Where(key+" = ?", value)
	}

	return query.Delete(&entity.Alamat{}).Error
}