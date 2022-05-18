package service

import "rpl-sixmath/model"

type AuthService interface {
	Login(request model.LoginRequest) (response model.LoginResponse, err error)
}
