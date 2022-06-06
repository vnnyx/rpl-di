package service

import (
	"errors"
	"rpl-sixmath/entity"
	"rpl-sixmath/model"
	"rpl-sixmath/repository/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

var (
	createdAt = time.Now()

	userEntity = entity.User{
		UserId:    1,
		Username:  "username",
		Email:     "email",
		Handphone: "08928273333",
		Password:  "password",
		Role:      "student",
		Gender:    "perempuan",
		Avatar:    "avatar",
		CreatedAt: createdAt,
	}
)

func TestRegisterUserSuccess(t *testing.T) {
	userRepoMock := new(mocks.UserRepository)

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(userEntity.Password), bcrypt.DefaultCost)

	student := entity.User{
		Username:  userEntity.Username,
		Handphone: userEntity.Handphone,
		Email:     userEntity.Email,
		Password:  string(passwordHash),
		Role:      userEntity.Role,
	}

	userRepoMock.On("InsertUser", student).Return(student, nil)

	userService := NewUserService(userRepoMock)

	request := model.StudentCreateRequest{
		Email:     userEntity.Email,
		Username:  userEntity.Username,
		Handphone: userEntity.Handphone,
		Password:  string(passwordHash),
	}

	response, err := userService.CreateStudent(request)

	assert.Nil(t, err)
	assert.Equal(t, userEntity.Email, response.Email)
	assert.Equal(t, userEntity.Handphone, response.Handphone)
	assert.Equal(t, userEntity.Username, response.Username)
}

func TestRegisterUserUsernameRegistered(t *testing.T) {
	userRepoMock := new(mocks.UserRepository)

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(userEntity.Password), bcrypt.DefaultCost)

	student := entity.User{
		Username:  userEntity.Username,
		Handphone: userEntity.Handphone,
		Email:     userEntity.Email,
		Password:  string(passwordHash),
		Role:      userEntity.Role,
	}

	userRepoMock.On("InsertUser", student).Return(student, errors.New("USERNAME_REGISTERED"))

	userService := NewUserService(userRepoMock)

	request := model.StudentCreateRequest{
		Email:     userEntity.Email,
		Username:  userEntity.Username,
		Handphone: userEntity.Handphone,
		Password:  string(passwordHash),
	}

	_, err := userService.CreateStudent(request)

	assert.Equal(t, errors.New("USERNAME_REGISTERED"), err)
}
