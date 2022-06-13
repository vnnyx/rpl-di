package service

import (
	"errors"
	"rpl-sixmath/entity"
	"rpl-sixmath/exception"
	"rpl-sixmath/model"
	"rpl-sixmath/repository"
	"rpl-sixmath/validation"
	"time"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (service *UserServiceImpl) CreateStudent(request model.StudentCreateRequest) (response model.StudentResponse, err error) {
	validation.Validate(request)

	student := entity.User{
		Username:  request.Username,
		Handphone: request.Handphone,
		Email:     request.Email,
		Password:  request.Password,
		Role:      "student",
		Avatar:    request.Avatar,
	}

	student, err = service.UserRepository.InsertUser(student)
	if err != nil {
		return model.StudentResponse{}, errors.New("USERNAME_REGISTERED")
	}
	response = model.StudentResponse{
		UserId:    student.UserId,
		Username:  student.Username,
		Email:     student.Email,
		Handphone: student.Handphone,
		Avatar:    student.Avatar,
		CreatedAt: student.CreatedAt,
	}

	return response, nil
}

func (service *UserServiceImpl) GetDataUser(month int) (response []model.GetUserResponse) {
	times := time.Now()
	for i := 0; i < month; i++ {
		date := times.AddDate(0, -i, 0)
		months := date.Format("01")
		year := date.Format("2006")
		total, err := service.UserRepository.GetDataUser(months, year)
		exception.PanicIfNeeded(err)
		response = append(response, model.GetUserResponse{
			Month:  date.Format("Jan"),
			Amount: total.Total,
		})
	}
	return response
}

func (service *UserServiceImpl) CreateTeacher(request model.TeacherCreateRequest) (response model.TeacherResponse, err error) {
	teacher := entity.User{
		Username:    request.Username,
		Email:       request.Email,
		Handphone:   request.Handphone,
		Password:    request.Password,
		Role:        "teacher",
		Certificate: request.Certificate,
		Avatar:      request.Avatar,
	}

	teacher, err = service.UserRepository.InsertUser(teacher)
	if err != nil {
		return model.TeacherResponse{}, errors.New("USERNAME_REGISTERED")
	}

	response = model.TeacherResponse{
		UserId:      teacher.UserId,
		Email:       teacher.Email,
		Username:    teacher.Username,
		Handphone:   teacher.Handphone,
		Certificate: teacher.Certificate,
		Avatar:      teacher.Avatar,
		CreatedAt:   teacher.CreatedAt,
	}

	return response, nil
}

func (service *UserServiceImpl) GetListTeacher() (response []model.TeacherResponse, err error) {
	users, err := service.UserRepository.FindAllTeacher()
	if err != nil {
		return []model.TeacherResponse{}, err
	}
	for _, user := range users {
		response = append(response, model.TeacherResponse{
			UserId:      user.UserId,
			Email:       user.Email,
			Username:    user.Username,
			Handphone:   user.Handphone,
			Certificate: user.Certificate,
			CreatedAt:   user.CreatedAt,
		})
	}
	return response, nil
}
