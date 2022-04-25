package service

import (
	"rpl-sixmath/entity"
	"rpl-sixmath/model"
	"rpl-sixmath/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: *userRepository}
}

func (service *UserServiceImpl) CreateStudent(request model.StudentCreateRequest) (response model.StudentCreateResponse) {
	student := entity.UserEntity{
		Username:  request.Username,
		Handphone: request.Handphone,
		Email:     request.Email,
		Password:  request.Password,
		Role:      "student",
	}

	service.UserRepository.InsertUser(student)
	response = model.StudentCreateResponse{
		Username:  student.Username,
		Email:     student.Email,
		Handphone: student.Handphone,
	}

	return response
}
