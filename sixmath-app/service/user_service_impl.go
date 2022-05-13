package service

import (
	"golang.org/x/crypto/bcrypt"
	"rpl-sixmath/entity"
	"rpl-sixmath/model"
	"rpl-sixmath/repository"
	"rpl-sixmath/validation"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: *userRepository}
}

func (service *UserServiceImpl) CreateStudent(request model.StudentCreateRequest) (response model.StudentCreateResponse) {
	validation.Validate(request)
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	student := entity.UserEntity{
		Username:  request.Username,
		Handphone: request.Handphone,
		Email:     request.Email,
		Password:  string(passwordHash),
		Role:      "student",
	}

	_, err = service.UserRepository.InsertUser(student)
	if err == nil {
		response = model.StudentCreateResponse{
			Username:  student.Username,
			Email:     student.Email,
			Handphone: student.Handphone,
		}
	}
	validation.ValidateUsername(err)

	return response
}
