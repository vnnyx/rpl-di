package repository

import (
	"gorm.io/gorm"
	"rpl-sixmath/entity"
	"rpl-sixmath/model"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: DB}
}

func (repo *UserRepositoryImpl) InsertUser(user entity.User) (entity.User, error) {
	err := repo.DB.Create(&user).Error
	return user, err
}

func (repo *UserRepositoryImpl) UpdateUser(user entity.User) (entity.User, error) {
	err := repo.DB.Where("user_id", user.UserId).Updates(&user).Error
	return user, err
}

func (repo *UserRepositoryImpl) DeleteUser(userId int) error {
	err := repo.DB.Where("user_id", userId).Delete(&entity.User{}).Error
	return err
}

func (repo *UserRepositoryImpl) FindUserById(userId int) (response entity.User, err error) {
	err = repo.DB.Where("user_id", userId).Find(&response).Error
	return response, err
}

func (repo *UserRepositoryImpl) FindUserByUsername(username string) (response entity.User, err error) {
	err = repo.DB.Where("username", username).Find(&response).Error
	return response, err
}

func (repo *UserRepositoryImpl) FindAllTeacher() (response []entity.User, err error) {
	err = repo.DB.Where("role", "teacher").Find(&response).Error
	return response, err
}

func (repo *UserRepositoryImpl) FindAll() (response []entity.User, err error) {
	err = repo.DB.Find(&response).Error
	return response, err
}

func (repo *UserRepositoryImpl) GetDataUser(month string, year string) (response model.GetTotalUser, err error) {
	err = repo.DB.Raw("SELECT COUNT(1) as total FROM users WHERE MONTH(created_at) = ? AND YEAR(created_at) = ?", month, year).Scan(&response).Error
	return response, err
}

func (repo *UserRepositoryImpl) FindUserByEmail(email string) (response entity.User, err error) {
	err = repo.DB.Where("email", email).First(&response).Error
	return response, err
}
