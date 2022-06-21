package service

import (
	"rpl-sixmath/entity"
	"rpl-sixmath/model"
)

type PasswordResetService interface {
	SendOtp(request model.SendOtpRequest) (response entity.User, err error)
	ValidateOtp(request model.SendOtpRequest) error
	PasswordReset(request model.PasswordResetRequest) error
}
