package repository

import (
	"rpl-sixmath/entity"
)

type UserRepository interface {
	InsertUser(user entity.User) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
	DeleteUser(userId int) error
	FindUserById(userId int) (response entity.User, err error)
	FindUserByUsername(username string) (response entity.User, err error)
	FindUserAll() (response []entity.User, err error)
}
