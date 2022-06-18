package service

import (
	"rpl-sixmath/model"
)

type UserService interface {
	CreateStudent(request model.StudentCreateRequest) (response model.StudentCreateResponse, err error)
	CreateTeacher(request model.TeacherCreateRequest) (response model.TeacherCreateResponse, err error)
	CreateParent(request model.ParentCreateRequest) (response model.ParentCreateResponse, err error)
	GetListTeacher() (response []model.TeacherCreateResponse, err error)
	GetDataUser(month int) (response []model.GetUserResponse)
}
