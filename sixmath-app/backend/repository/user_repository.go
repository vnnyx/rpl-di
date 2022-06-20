package repository

import (
	"rpl-sixmath/entity"
	"rpl-sixmath/model"
)

type UserRepository interface {
	InsertUser(user entity.User) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
	DeleteUser(userId int) error
	FindUserById(userId int) (response entity.User, err error)
	FindUserByEmail(email string) (response entity.User, err error)
	FindUserByUsername(username string) (response entity.User, err error)
	FindAllTeacher() (response []entity.User, err error)
	FindAll() (response []entity.User, err error)
	GetDataUser(month string, year string) (response model.GetTotalUser, err error)
}
