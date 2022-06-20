package repository

import (
	"rpl-sixmath/entity"
)

type PasswordResetRepository interface {
	InsertToken(reset entity.PasswordReset) (entity.PasswordReset, error)
	ValidateToken(otp int, email string) (response entity.PasswordReset, err error)
	DeleteToken(email string) error
}
