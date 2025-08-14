package repository

import (
	"svc-llt-golang/domain/llt"
	"svc-llt-golang/entity"
	"svc-llt-golang/valueobject"

	"gorm.io/gorm"
)

type mysqlLltRepository struct {
	db *gorm.DB
}

func NewMysqlLltRepository(db *gorm.DB) llt.Repository {
	return &mysqlLltRepository{db}
}

func (r *mysqlLltRepository) GetAllElderlyCare(param map[string]interface{}) ([]valueobject.ElderlyCare, error) {
	var elderlyCareList []entity.ElderlyCare
	query := r.db
	
	if err := query.Find(&elderlyCareList).Error; err != nil {
		return nil, err
	}
	
	var result []valueobject.ElderlyCare
	for _, elderlyCare := range elderlyCareList {
		result = append(result, valueobject.ElderlyCare{ElderlyCare: elderlyCare})
	}
	
	return result, nil
}

func (r *mysqlLltRepository) GetOneElderlyCare(param map[string]interface{}) (valueobject.ElderlyCare, error) {
	var elderlyCare entity.ElderlyCare
	query := r.db
	
	if uuid, ok := param["uuid"].(string); ok {
		query = query.Where("uuid = ?", uuid)
	}
	
	if err := query.First(&elderlyCare).Error; err != nil {
		return valueobject.ElderlyCare{}, err
	}
	
	return valueobject.ElderlyCare{ElderlyCare: elderlyCare}, nil
}

func (r *mysqlLltRepository) CreateElderlyCare(elderlyCare valueobject.ElderlyCare) error {
	return r.db.Create(&elderlyCare.ElderlyCare).Error
}

func (r *mysqlLltRepository) UpdateElderlyCare(param map[string]interface{}, data map[string]interface{}) error {
	query := r.db.Model(&entity.ElderlyCare{})
	
	for key, value := range param {
		query = query.Where(key+" = ?", value)
	}
	
	return query.Updates(data).Error
}

func (r *mysqlLltRepository) DeleteElderlyCare(param map[string]interface{}) error {
	query := r.db.Model(&entity.ElderlyCare{})
	
	for key, value := range param {
		query = query.Where(key+" = ?", value)
	}
	
	return query.Delete(&entity.ElderlyCare{}).Error
}