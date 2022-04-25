package repository

import "rpl-sixmath/entity"

type UserRepository interface {
	InsertUser(user entity.UserEntity) (entity.UserEntity, error)
	UpdateUser(user entity.UserEntity) (entity.UserEntity, error)
	DeleteUser(userId int) error
	FindUserById(userId int) (response entity.UserEntity, err error)
	FindUserByUsername(username string) (response entity.UserEntity, err error)
	FindUserAll() (response []entity.UserEntity, err error)
}
