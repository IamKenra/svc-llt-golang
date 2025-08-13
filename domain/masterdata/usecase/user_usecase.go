package usecase

import (
	"errors"
	"time"

	"svc-llt-golang/domain/masterdata"
	"svc-llt-golang/utils/utils"
	"svc-llt-golang/valueobject"

	"github.com/golang-jwt/jwt/v5"
)

type userUsecase struct {
	userRepo masterdata.Repository
	jwtKey   string
}

func NewUserUsecase(userRepo masterdata.Repository, jwtKey string) masterdata.Usecase {
	return &userUsecase{userRepo: userRepo, jwtKey: jwtKey}
}

func (u *userUsecase) Login(username, password string) (string, error) {
	user, err := u.userRepo.FindByUsername(username)
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

	tokenString, err := token.SignedString([]byte(u.jwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (u *userUsecase) GetAllUser(param map[string]interface{}) ([]valueobject.User, error) {
	return u.userRepo.GetAllUser(param)
}

func (u *userUsecase) GetOneUser(param map[string]interface{}) (valueobject.User, error) {
	return u.userRepo.GetOneUser(param)
}

func (u *userUsecase) CreateUser(user valueobject.User) error {
	return u.userRepo.CreateUser(user)
}

func (u *userUsecase) UpdateUser(param map[string]interface{}, data map[string]interface{}) error {
	return u.userRepo.UpdateUser(param, data)
}

func (u *userUsecase) DeleteUser(param map[string]interface{}) error {
	return u.userRepo.DeleteUser(param)
}