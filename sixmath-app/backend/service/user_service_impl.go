package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
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

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: *userRepository}
}

func (service *UserServiceImpl) CreateStudent(request model.StudentCreateRequest) (response model.StudentCreateResponse) {
	validation.Validate(request)

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	student := entity.User{
		Username:  request.Username,
		Handphone: request.Handphone,
		Email:     request.Email,
		Password:  string(passwordHash),
		Role:      "student",
	}

	student, err = service.UserRepository.InsertUser(student)
	fmt.Println(student)
	if err != nil {
		exception.PanicIfNeeded("USERNAME_REGISTERED")
	}
	response = model.StudentCreateResponse{
		UserId:    student.UserId,
		Username:  student.Username,
		Email:     student.Email,
		Handphone: student.Handphone,
		CreatedAt: student.CreatedAt,
	}

	return response
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
