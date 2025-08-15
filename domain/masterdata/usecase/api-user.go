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

func (masterdata masterdataUsecase) Register(payload valueobject.UserRegisterRequest) (valueobject.UserRegisterResponse, error) {
	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return valueobject.UserRegisterResponse{}, errors.New("failed to hash password")
	}

	err = masterdata.repository.CreateAuth(payload.Username, hashedPassword)
	if err != nil {
		return valueobject.UserRegisterResponse{}, errors.New("username already exists")
	}

	userUUID := uuid.New().String()
	err = masterdata.repository.CreateUser(userUUID, payload.Nama, payload.Email)
	if err != nil {
		return valueobject.UserRegisterResponse{}, errors.New("email already exists")
	}

	return valueobject.UserRegisterResponse{
		Message: "User registered successfully",
		UUID:    userUUID,
	}, nil
}

func (masterdata masterdataUsecase) GetAllUser(param map[string]interface{}) ([]valueobject.User, error) {
	return masterdata.repository.GetAllUser(param)
}

func (masterdata masterdataUsecase) GetOneUser(param map[string]interface{}) (valueobject.User, error) {
	return masterdata.repository.GetOneUser(param)
}

func (masterdata masterdataUsecase) StoreUser(payload valueobject.UserPayloadInsert) (valueobject.UserPayloadInsert, error) {
	userUUID := uuid.New().String()
	err := masterdata.repository.CreateUser(userUUID, payload.Data[0].Nama, payload.Data[0].Email)
	return payload, err
}

func (masterdata masterdataUsecase) UpdateUser(payload valueobject.UserPayloadUpdate) error {
	param := map[string]interface{}{
		"uuid": payload.Data.Param.UUID,
	}
	data := map[string]interface{}{
		"nama":  payload.Data.Body.Nama,
		"email": payload.Data.Body.Email,
	}
	return masterdata.repository.UpdateUser(param, data)
}

func (masterdata masterdataUsecase) DeleteUser(payload valueobject.UserPayloadDelete) error {
	for _, user := range payload.Param {
		param := map[string]interface{}{
			"uuid": user.UUID,
		}
		err := masterdata.repository.DeleteUser(param)
		if err != nil {
			return err
		}
	}
	return nil
}
