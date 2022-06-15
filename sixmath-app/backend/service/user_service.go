package service

import (
	"rpl-sixmath/model"
)

type UserService interface {
	CreateStudent(request model.StudentCreateRequest) (response model.StudentResponse, err error)
	CreateTeacher(request model.TeacherCreateRequest) (response model.TeacherResponse, err error)
	GetListTeacher() (response []model.TeacherResponse, err error)
	GetDataUser(month int) (response []model.GetUserResponse)
}
