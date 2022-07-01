package service

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"html/template"
	"math/rand"
	"rpl-sixmath/entity"
	"rpl-sixmath/exception"
	"rpl-sixmath/helper"
	"rpl-sixmath/model"
	"rpl-sixmath/repository"
	"time"
)

//go:embed templates/*.gohtml
var templates embed.FS

type PasswordResetServiceImpl struct {
	repository.PasswordResetRepository
	repository.UserRepository
}

func NewPasswordResetService(passwordResetRepository repository.PasswordResetRepository, userRepository repository.UserRepository) PasswordResetService {
	return &PasswordResetServiceImpl{
		PasswordResetRepository: passwordResetRepository,
		UserRepository:          userRepository}
}

func (service *PasswordResetServiceImpl) SendOtp(request model.SendOtpRequest) (response entity.User, err error) {
	user, err := service.UserRepository.FindUserByEmail(request.Email)
	if err != nil {
		return entity.User{}, errors.New("EMAIL_NOT_FOUND")
	}
	fmt.Println(user)
	otp := rand.Intn(9999-1000) + 1000
	passwordReset := entity.PasswordReset{
		UserEmail: request.Email,
		Token:     otp,
		ExpiredAt: time.Now().Add(time.Minute * 5),
	}
	passwordReset, err = service.PasswordResetRepository.InsertToken(passwordReset)
	if err != nil {
		return entity.User{}, err
	}

	t, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		return entity.User{}, err
	}
	buff := new(bytes.Buffer)
	err = t.ExecuteTemplate(buff, "send_otp.gohtml", map[string]interface{}{
		"Username": user.Username,
		"Otp":      otp,
	})
	if err != nil {
		return entity.User{}, err
	}
	err = helper.SendEmailTo(user.Email, buff)
	if err != nil {
		return entity.User{}, err
	}
	response = entity.User{
		UserId:   user.UserId,
		Username: user.Username,
		Email:    user.Email,
	}
	return response, nil
}

func (service *PasswordResetServiceImpl) ValidateOtp(request model.SendOtpRequest) error {
	validate, err := service.PasswordResetRepository.ValidateToken(request.Otp, request.Email)
	if err == gorm.ErrRecordNotFound {
		return errors.New("TOKEN_INVALID")
	}
	err = service.PasswordResetRepository.DeleteToken(validate.UserEmail)
	exception.PanicIfNeeded(err)
	return nil
}

func (service *PasswordResetServiceImpl) PasswordReset(request model.PasswordResetRequest) error {
	user, err := service.UserRepository.FindUserById(request.UserId)
	if err != nil {
		return errors.New("USER_NOT_FOUND")
	}
	_, err = service.UserRepository.UpdateUser(entity.User{
		UserId:          request.UserId,
		Username:        user.Username,
		Email:           user.Email,
		Handphone:       user.Handphone,
		Password:        request.Password,
		Role:            user.Role,
		Avatar:          user.Avatar,
		Certificate:     user.Certificate,
		StudentUsername: user.StudentUsername,
		Age:             user.Age,
		Description:     user.Description,
		CreatedAt:       user.CreatedAt,
	})
	exception.PanicIfNeeded(err)
	return nil
}
