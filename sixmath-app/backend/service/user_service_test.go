package service

import (
	"fmt"
	"rpl-sixmath/entity"
	"rpl-sixmath/repository/mocks"
	"testing"
)

func TestRegisterUserSuccess(t *testing.T) {
	userRepoMock := new(mocks.UserRepository)

	student := entity.User{
		Username:  "usertest",
		Handphone: "081234567890",
		Email:     "test@email.com",
		Password:  "password",
		Role:      "student",
	}

	fmt.Println(userRepoMock)
	fmt.Println(student)
}
