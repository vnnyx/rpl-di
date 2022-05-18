package service

import (
	"rpl-sixmath/model"
)

type UserService interface {
	CreateStudent(request model.StudentCreateRequest) (response model.StudentCreateResponse)
	GetDataUser(month int) (response []model.GetUserResponse)
}
