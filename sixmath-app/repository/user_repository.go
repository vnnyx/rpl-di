package repository

import "rpl-sixmath/entity"

type UserRepository interface {
	InsertUser(user entity.UserEntity) (entity.UserEntity, error)
	UpdateUser(user entity.UserEntity) (entity.UserEntity, error)
	DeleteUser(userId int) error
	FindUserById(userId int) (entity.UserEntity, error)
	FindUserByUsername(username string) (entity.UserEntity, error)
	FindUserAll() ([]entity.UserEntity, error)
}
