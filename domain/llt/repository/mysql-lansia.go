package repository

import (
	"time"
	"svc-llt-golang/domain/llt"
	"svc-llt-golang/entity"
	"svc-llt-golang/valueobject"
	"svc-llt-golang/utils/utils"

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
	
	// Apply filters if provided
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

func (db *mysqlLltRepository) CreateLansia(uuid, name string, age int, status string) error {
	randomID, err := utils.GenerateRandomID()
	if err != nil {
		return err
	}

	lansia := entity.Lansia{
		ID:        randomID,
		UUID:      uuid,
		Name:      name,
		Age:       age,
		Status:    status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return db.db.Create(&lansia).Error
}

func (db *mysqlLltRepository) UpdateLansia(param map[string]interface{}, data map[string]interface{}) error {
	query := db.db.Model(&entity.Lansia{})
	
	for key, value := range param {
		query = query.Where(key+" = ?", value)
	}
	
	// Add updated timestamp
	data["updated_at"] = time.Now()
	
	return query.Updates(data).Error
}

func (db *mysqlLltRepository) DeleteLansia(param map[string]interface{}) error {
	query := db.db.Model(&entity.Lansia{})
	
	for key, value := range param {
		query = query.Where(key+" = ?", value)
	}
	
	return query.Delete(&entity.Lansia{}).Error
}