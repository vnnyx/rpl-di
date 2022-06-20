package repository

import (
	"gorm.io/gorm"
	"rpl-sixmath/entity"
	"time"
)

type PasswordResetRepositoryImpl struct {
	*gorm.DB
}

func NewPasswordResetRepository(DB *gorm.DB) PasswordResetRepository {
	return &PasswordResetRepositoryImpl{DB: DB}
}

func (repo *PasswordResetRepositoryImpl) InsertToken(reset entity.PasswordReset) (entity.PasswordReset, error) {
	err := repo.DB.Create(&reset).Error
	return reset, err
}

func (repo *PasswordResetRepositoryImpl) ValidateToken(otp int, email string) (response entity.PasswordReset, err error) {
	err = repo.DB.First(&response).Where("token", otp).Where("expired_at >", time.Now()).Where("user_email", email).Error
	return response, err
}

func (repo *PasswordResetRepositoryImpl) DeleteToken(email string) error {
	err := repo.DB.Where("user_email", email).Delete(&entity.PasswordReset{}).Error
	return err
}
