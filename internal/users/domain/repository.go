package domain

import "svc-llt-golang/internal/users/domain/entity"

type UserRepository interface {
	FindByUsername(username string) (*entity.User, error)
	FindByUUID(uuid string) (*entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(id uint) error
}