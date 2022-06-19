package service

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"rpl-sixmath/entity"
	"rpl-sixmath/exception"
	"rpl-sixmath/helper"
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

func (service *UserServiceImpl) CreateStudent(request model.StudentCreateRequest) (response model.StudentCreateResponse, err error) {
	validation.CreateStudentValidation(request)

	user, err := service.UserRepository.FindUserByUsername(request.Username)
	if (user != entity.User{}) {
		return model.StudentCreateResponse{}, errors.New("USERNAME_REGISTERED")
	}

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
		return model.StudentCreateResponse{}, err
	}

	td := helper.CreateToken(model.JwtPayload{
		UserID:   student.UserId,
		Avatar:   student.Avatar,
		Username: student.Username,
		Email:    student.Email,
		Role:     student.Role,
	})

	response = model.StudentCreateResponse{
		AccessToken: td.AccessToken,
		UserId:      student.UserId,
		Username:    student.Username,
		Email:       student.Email,
		Handphone:   student.Handphone,
		Role:        student.Role,
		Avatar:      request.Avatar,
		CreatedAt:   student.CreatedAt,
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

func (service *UserServiceImpl) CreateTeacher(request model.TeacherCreateRequest) (response model.TeacherCreateResponse, err error) {
	validation.CreateTeacherValidation(request)

	user, err := service.UserRepository.FindUserByUsername(request.Username)
	if (user != entity.User{}) {
		return model.TeacherCreateResponse{}, errors.New("USERNAME_REGISTERED")
	}

	var ctx *fiber.Ctx
	resAvatar, err := helper.UploadToCloudinary(ctx, request.Avatar, ".:/sixmath/avatar/", request.Username+"_avatar")
	exception.PanicIfNeeded(err)
	resCertificate, err := helper.UploadToCloudinary(ctx, request.Certificate, ".:/sixmath/certificate/", request.Username+"_certificate")
	exception.PanicIfNeeded(err)
	AvatarUrl := resAvatar.SecureURL
	CertificateUrl := resCertificate.SecureURL
	teacher := entity.User{
		Username:    request.Username,
		Email:       request.Email,
		Handphone:   request.Handphone,
		Password:    request.Password,
		Role:        "teacher",
		Certificate: CertificateUrl,
		Avatar:      AvatarUrl,
		Age:         request.Age,
		Description: request.Description,
	}

	teacher, err = service.UserRepository.InsertUser(teacher)
	if err != nil {
		return model.TeacherCreateResponse{}, err
	}

	td := helper.CreateToken(model.JwtPayload{
		UserID:   teacher.UserId,
		Avatar:   teacher.Avatar,
		Username: teacher.Username,
		Email:    teacher.Email,
		Role:     teacher.Role,
	})

	response = model.TeacherCreateResponse{
		AccessToken: td.AccessToken,
		UserId:      teacher.UserId,
		Email:       teacher.Email,
		Username:    teacher.Username,
		Handphone:   teacher.Handphone,
		Role:        teacher.Role,
		Certificate: teacher.Certificate,
		Avatar:      teacher.Avatar,
		Age:         teacher.Age,
		Description: teacher.Description,
		CreatedAt:   teacher.CreatedAt,
	}

	return response, nil
}

func (service *UserServiceImpl) CreateParent(request model.ParentCreateRequest) (response model.ParentCreateResponse, err error) {
	validation.CreateParentValidation(request)

	user, err := service.UserRepository.FindUserByUsername(request.Username)
	if (user != entity.User{}) {
		return model.ParentCreateResponse{}, errors.New("USERNAME_REGISTERED")
	}
	student, err := service.UserRepository.FindUserByUsername(request.StudentUsername)
	if (student == entity.User{}) {
		return model.ParentCreateResponse{}, errors.New("STUDENT_NOT_FOUND")
	}
	var ctx *fiber.Ctx
	resAvatar, err := helper.UploadToCloudinary(ctx, request.Avatar, ".:/sixmath/avatar/", request.Username+"_avatar")
	exception.PanicIfNeeded(err)
	AvatarUrl := resAvatar.SecureURL

	parent := entity.User{
		Username:        request.Username,
		Email:           request.Email,
		Handphone:       request.Handphone,
		Password:        request.Password,
		Role:            "parent",
		Avatar:          AvatarUrl,
		StudentUsername: request.StudentUsername,
	}

	parent, err = service.UserRepository.InsertUser(parent)
	if err != nil {
		return model.ParentCreateResponse{}, err
	}

	td := helper.CreateToken(model.JwtPayload{
		UserID:   parent.UserId,
		Avatar:   parent.Avatar,
		Username: parent.Username,
		Email:    parent.Email,
		Role:     parent.Role,
	})

	response = model.ParentCreateResponse{
		AccessToken:     td.AccessToken,
		Email:           parent.Email,
		Username:        parent.Username,
		Handphone:       parent.Handphone,
		StudentUsername: parent.StudentUsername,
		Role:            parent.Role,
		Avatar:          parent.Avatar,
	}

	return response, nil
}

func (service *UserServiceImpl) GetListTeacher() (response []model.TeacherCreateResponse, err error) {
	users, err := service.UserRepository.FindAllTeacher()
	if err != nil {
		return []model.TeacherCreateResponse{}, err
	}
	for _, teacher := range users {
		response = append(response, model.TeacherCreateResponse{
			UserId:      teacher.UserId,
			Email:       teacher.Email,
			Username:    teacher.Username,
			Handphone:   teacher.Handphone,
			Certificate: teacher.Certificate,
			Avatar:      teacher.Avatar,
			Age:         teacher.Age,
			Description: teacher.Description,
			CreatedAt:   teacher.CreatedAt,
		})
	}
	return response, nil
}
