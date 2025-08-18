package repository

import (
	"svc-llt-golang/entity"
	"svc-llt-golang/utils/utils"
	"svc-llt-golang/valueobject"
	"time"
)

func (db *mysqlLltRepository) GetAllIdentitas(param map[string]interface{}) ([]valueobject.Identitas, error) {
	var identitasList []entity.Identitas
	query := db.db

	for key, value := range param {
		query = query.Where(key+" = ?", value)
	}

	if err := query.Find(&identitasList).Error; err != nil {
		return nil, err
	}

	var result []valueobject.Identitas
	for _, identitas := range identitasList {
		result = append(result, valueobject.IdentitasFromEntity(identitas))
	}

	return result, nil
}

func (db *mysqlLltRepository) GetOneIdentitas(param map[string]interface{}) (valueobject.Identitas, error) {
	var identitas entity.Identitas
	query := db.db

	for key, value := range param {
		query = query.Where(key+" = ?", value)
	}

	if err := query.First(&identitas).Error; err != nil {
		return valueobject.Identitas{}, err
	}

	return valueobject.IdentitasFromEntity(identitas), nil
}

func (db *mysqlLltRepository) CreateIdentitas(uuid, nik, namaDepan string, namaBelakang *string, userInput string) error {
	randomID, err := utils.GenerateRandomID()
	if err != nil {
		return err
	}

	identitas := entity.Identitas{
		ID:        uint64(randomID),
		UUID:      uuid,
		Nik:       nik,
		NamaDepan: namaDepan,
		FlagAktif: true,
		TglInput:  time.Now(),
		TglUpdate: time.Now(),
	}

	if namaBelakang != nil {
		identitas.NamaBelakang = namaBelakang
	}
	if userInput != "" {
		identitas.UserInput = &userInput
	}

	return db.db.Create(&identitas).Error
}

func (db *mysqlLltRepository) UpdateIdentitas(param map[string]interface{}, data map[string]interface{}) error {
	query := db.db.Model(&entity.Identitas{})

	for key, value := range param {
		query = query.Where(key+" = ?", value)
	}

	data["tgl_update"] = time.Now()
	return query.Updates(data).Error
}

func (db *mysqlLltRepository) DeleteIdentitas(param map[string]interface{}) error {
	query := db.db.Model(&entity.Identitas{})

	for key, value := range param {
		query = query.Where(key+" = ?", value)
	}

	return query.Delete(&entity.Identitas{}).Error
}