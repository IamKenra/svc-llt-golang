package usecase

import (
	"errors"
	"time"

	"svc-llt-golang/utils/utils"
	"svc-llt-golang/valueobject"

	"github.com/golang-jwt/jwt/v5"
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

func (masterdata masterdataUsecase) GetAllUser(param map[string]interface{}) ([]valueobject.User, error) {
	return masterdata.repository.GetAllUser(param)
}

func (masterdata masterdataUsecase) GetOneUser(param map[string]interface{}) (valueobject.User, error) {
	return masterdata.repository.GetOneUser(param)
}

func (masterdata masterdataUsecase) StoreUser(payload valueobject.UserPayloadInsert) (valueobject.UserPayloadInsert, error) {
	err := masterdata.repository.CreateUser(payload.Data[0])
	return payload, err
}

func (masterdata masterdataUsecase) UpdateUser(payload valueobject.UserPayloadUpdate) error {
	param := map[string]interface{}{
		"uuid": payload.Data.Param.UUID,
	}
	data := map[string]interface{}{
		"username": payload.Data.Body.Username,
		"password": payload.Data.Body.Password,
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
