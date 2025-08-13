package usecase

import (
	"errors"
	"time"

	"svc-llt-golang/internal/users/domain"
	"svc-llt-golang/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
)

type UserUsecase interface {
	Login(username, password string) (string, error)
}

type userUsecase struct {
	userRepo domain.UserRepository
	jwtKey   string
}

func NewUserUsecase(userRepo domain.UserRepository, jwtKey string) UserUsecase {
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