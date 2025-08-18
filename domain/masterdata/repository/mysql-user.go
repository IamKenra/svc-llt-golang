package repository

import (
	"time"
	"svc-llt-golang/domain/masterdata"
	"svc-llt-golang/entity"
	"svc-llt-golang/valueobject"
	"svc-llt-golang/utils/utils"

	"gorm.io/gorm"
)

type mysqlMasterdataRepository struct {
	db *gorm.DB
}

func NewMysqlMasterdataRepository(db *gorm.DB) masterdata.Repository {
	return &mysqlMasterdataRepository{db}
}

func (db *mysqlMasterdataRepository) GetAllUser(param map[string]interface{}) ([]valueobject.User, error) {
	var users []entity.User
	query := db.db
	
	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	
	var result []valueobject.User
	for _, user := range users {
		result = append(result, valueobject.UserFromEntity(user))
	}
	
	return result, nil
}

func (db *mysqlMasterdataRepository) GetOneUser(param map[string]interface{}) (valueobject.User, error) {
	var user entity.User
	query := db.db
	
	if uuid, ok := param["uuid"].(string); ok {
		query = query.Where("uuid = ?", uuid)
	}
	
	if err := query.First(&user).Error; err != nil {
		return valueobject.User{}, err
	}
	
	return valueobject.UserFromEntity(user), nil
}

func (db *mysqlMasterdataRepository) FindByUsername(username string) (valueobject.Auth, error) {
	var auth entity.Auth
	if err := db.db.Where("username = ?", username).First(&auth).Error; err != nil {
		return valueobject.Auth{}, err
	}
	return valueobject.AuthFromEntity(auth), nil
}

func (db *mysqlMasterdataRepository) FindByUUID(uuid string) (valueobject.User, error) {
	var user entity.User
	if err := db.db.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		return valueobject.User{}, err
	}
	return valueobject.UserFromEntity(user), nil
}

func (db *mysqlMasterdataRepository) CreateAuth(username, password string) error {
	randomID, err := utils.GenerateRandomID()
	if err != nil {
		return err
	}

	auth := entity.Auth{
		ID:        uint64(randomID),
		Username:  username,
		Password:  password,
		TglInput:  time.Now(),
		TglUpdate: time.Now(),
		UserInput: "system",
	}
	return db.db.Create(&auth).Error
}

func (db *mysqlMasterdataRepository) CreateUser(uuid, nama, email string) error {
	randomID, err := utils.GenerateRandomID()
	if err != nil {
		return err
	}

	user := entity.User{
		ID:        uint64(randomID),
		UUID:      uuid,
		Nama:      nama,
		Email:     email,
		TglInput:  time.Now(),
		TglUpdate: time.Now(),
		UserInput: "system",
	}
	return db.db.Create(&user).Error
}

func (db *mysqlMasterdataRepository) UpdateUser(param map[string]interface{}, data map[string]interface{}) error {
	query := db.db.Model(&entity.User{})
	
	for key, value := range param {
		query = query.Where(key+" = ?", value)
	}
	
	return query.Updates(data).Error
}

func (db *mysqlMasterdataRepository) DeleteUser(param map[string]interface{}) error {
	query := db.db.Model(&entity.User{})
	
	for key, value := range param {
		query = query.Where(key+" = ?", value)
	}
	
	return query.Delete(&entity.User{}).Error
}

