package repository

import (
	"svc-llt-golang/domain/masterdata"
	"svc-llt-golang/entity"
	"svc-llt-golang/valueobject"

	"gorm.io/gorm"
)

type mysqlMasterdataRepository struct {
	db *gorm.DB
}

func NewMysqlMasterdataRepository(db *gorm.DB) masterdata.Repository {
	return &mysqlMasterdataRepository{db}
}

func (r *mysqlMasterdataRepository) GetAllUser(param map[string]interface{}) ([]valueobject.User, error) {
	var users []entity.User
	query := r.db
	
	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	
	var result []valueobject.User
	for _, user := range users {
		result = append(result, valueobject.User{User: user})
	}
	
	return result, nil
}

func (r *mysqlMasterdataRepository) GetOneUser(param map[string]interface{}) (valueobject.User, error) {
	var user entity.User
	query := r.db
	
	if uuid, ok := param["uuid"].(string); ok {
		query = query.Where("uuid = ?", uuid)
	}
	
	if err := query.First(&user).Error; err != nil {
		return valueobject.User{}, err
	}
	
	return valueobject.User{User: user}, nil
}

func (r *mysqlMasterdataRepository) FindByUsername(username string) (valueobject.User, error) {
	var user entity.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return valueobject.User{}, err
	}
	return valueobject.User{User: user}, nil
}

func (r *mysqlMasterdataRepository) FindByUUID(uuid string) (valueobject.User, error) {
	var user entity.User
	if err := r.db.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		return valueobject.User{}, err
	}
	return valueobject.User{User: user}, nil
}

func (r *mysqlMasterdataRepository) CreateUser(user valueobject.User) error {
	return r.db.Create(&user.User).Error
}

func (r *mysqlMasterdataRepository) UpdateUser(param map[string]interface{}, data map[string]interface{}) error {
	query := r.db.Model(&entity.User{})
	
	for key, value := range param {
		query = query.Where(key+" = ?", value)
	}
	
	return query.Updates(data).Error
}

func (r *mysqlMasterdataRepository) DeleteUser(param map[string]interface{}) error {
	query := r.db.Model(&entity.User{})
	
	for key, value := range param {
		query = query.Where(key+" = ?", value)
	}
	
	return query.Delete(&entity.User{}).Error
}