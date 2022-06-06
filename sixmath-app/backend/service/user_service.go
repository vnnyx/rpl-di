package service

import (
	"rpl-sixmath/model"
)

type UserService interface {
	CreateStudent(request model.StudentCreateRequest) (response model.StudentCreateResponse, err error)
	GetDataUser(month int) (response []model.GetUserResponse)
}
