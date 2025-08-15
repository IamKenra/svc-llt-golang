package usecase

import (
	"errors"
	"time"

	"svc-llt-golang/utils/utils"
	"svc-llt-golang/valueobject"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (masterdata masterdataUsecase) Login(username, password string) (string, error) {
	user, err := masterdata.repository.FindByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid": user.UUID,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(masterdata.jwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (masterdata masterdataUsecase) Register(payload valueobject.UserRegisterRequest) (string, error) {
	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return "", errors.New("failed to hash password")
	}

	err = masterdata.repository.CreateAuth(payload.Username, hashedPassword)
	if err != nil {
		return "", errors.New("username already exists")
	}

	userUUID := uuid.New().String()
	err = masterdata.repository.CreateUser(userUUID, payload.Nama, payload.Email)
	if err != nil {
		return "", errors.New("email already exists")
	}

	return userUUID, nil
}

func (masterdata masterdataUsecase) GetAllUser(param map[string]interface{}) ([]valueobject.User, error) {
	return masterdata.repository.GetAllUser(param)
}

func (masterdata masterdataUsecase) GetOneUser(param map[string]interface{}) (valueobject.User, error) {
	return masterdata.repository.GetOneUser(param)
}

func (masterdata masterdataUsecase) StoreUser(payload valueobject.UserPayloadInsert) (valueobject.UserPayloadInsert, error) {
	for i := range payload.Data {
		payload.Data[i].UUID = uuid.New().String()
	}

	err := masterdata.ProcessStoreUser(payload)
	return payload, err
}

func (masterdata masterdataUsecase) UpdateUser(payload valueobject.UserPayloadUpdate) error {
	return masterdata.ProcessUpdateUser(payload)
}

func (masterdata masterdataUsecase) DeleteUser(payload valueobject.UserPayloadDelete) error {
	return masterdata.ProcessDeleteUser(payload)
}
