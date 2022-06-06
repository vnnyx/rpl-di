package service

import (
	"errors"
	"fmt"
	"rpl-sixmath/helper"
	"rpl-sixmath/model"
	"rpl-sixmath/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) AuthService {
	return &AuthServiceImpl{UserRepository: *userRepository}
}

func (service *AuthServiceImpl) Login(request model.LoginRequest) (response model.LoginResponse, err error) {
	user, err := service.UserRepository.FindUserByUsername(request.UserName)
	if err != nil {
		return response, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return response, errors.New(model.UNAUTHORIZATION)
	}

	td := helper.CreateToken(model.JwtPayload{
		UserID:   user.UserId,
		Avatar:   user.Avatar,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	})

	response = model.LoginResponse{
		AccessToken: td.AccessToken,
		UserID:      user.UserId,
		Email:       user.Email,
		Avatar:      user.Avatar,
		Username:    user.Username,
		Role:        user.Role,
	}

	return response, nil
}
