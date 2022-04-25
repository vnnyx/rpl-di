package repository

import (
	"gorm.io/gorm"
	"rpl-sixmath/entity"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: DB}
}

func (repo *UserRepositoryImpl) InsertUser(user entity.UserEntity) (entity.UserEntity, error) {
	err := repo.DB.Create(&user).Error
	return user, err
}

func (repo *UserRepositoryImpl) UpdateUser(user entity.UserEntity) (entity.UserEntity, error) {
	err := repo.DB.Where("user_id", user.UserId).Updates(&user).Error
	return user, err
}

func (repo *UserRepositoryImpl) DeleteUser(userId int) error {
	err := repo.DB.Where("user_id", userId).Delete(&entity.UserEntity{}).Error
	return err
}

func (repo *UserRepositoryImpl) FindUserById(userId int) (entity.UserEntity, error) {
	err := repo.DB.Where("user_id", userId).Find(&entity.UserEntity{}).Error
	return entity.UserEntity{}, err
}

func (repo *UserRepositoryImpl) FindUserByUsername(username string) (entity.UserEntity, error) {
	err := repo.DB.Where("username", username).Find(&entity.UserEntity{}).Error
	return entity.UserEntity{}, err
}

func (repo *UserRepositoryImpl) FindUserAll() ([]entity.UserEntity, error) {
	err := repo.DB.Find(&entity.UserEntity{}).Error
	return []entity.UserEntity{}, err
}
