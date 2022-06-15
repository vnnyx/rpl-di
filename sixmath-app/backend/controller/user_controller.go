package controller

import (
	"rpl-sixmath/exception"
	"rpl-sixmath/helper"
	"rpl-sixmath/middleware"
	"rpl-sixmath/model"
	"rpl-sixmath/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(service *service.UserService) UserController {
	return UserController{UserService: *service}
}

func (controller *UserController) Route(app *fiber.App) {
	router := app.Group("/api/user")
	router.Post("/student", controller.CreateStudent)
	router.Post("/teacher", controller.CreateTeacher)
	router.Get("/teacher/all", middleware.CheckToken(), controller.GetAllTeacher)
	app.Get("/api/dashboard/statistik-pendaftaran", middleware.CheckToken(), controller.GetDataUser)
}

func (controller *UserController) CreateStudent(c *fiber.Ctx) error {
	var request model.StudentCreateRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	avatar, err := c.FormFile("avatar")
	exception.PanicIfNeeded(err)

	resAvatar, err := helper.UploadToCloudinary(c, avatar, "./uploads/image:/sixmath/avatar", request.Username+"_avatar")
	exception.PanicIfNeeded(err)
	avatarUrl := resAvatar.SecureURL
	request.Avatar = avatarUrl

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	request.Password = string(passwordHash)

	response, err := controller.UserService.CreateStudent(request)
	exception.PanicIfNeeded(err)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *UserController) CreateTeacher(c *fiber.Ctx) error {
	var request model.TeacherCreateRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	certificate, _ := c.FormFile("certificate")
	avatar, _ := c.FormFile("avatar")
	resCertificate, err := helper.UploadToCloudinary(c, certificate, "./uploads/image:/sixmath/certificate", request.Username+"_certificate")
	exception.PanicIfNeeded(err)
	resAvatar, err := helper.UploadToCloudinary(c, avatar, "./uploads/image:/sixmath/avatar", request.Username+"_avatar")
	exception.PanicIfNeeded(err)

	certificateUrl := resCertificate.SecureURL
	AvatarUrl := resAvatar.SecureURL
	request.Certificate = certificateUrl
	request.Avatar = AvatarUrl

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	request.Password = string(passwordHash)

	response, err := controller.UserService.CreateTeacher(request)
	exception.PanicIfNeeded(err)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *UserController) GetAllTeacher(c *fiber.Ctx) error {
	response, err := controller.UserService.GetListTeacher()
	exception.PanicIfNeeded(err)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *UserController) GetDataUser(c *fiber.Ctx) error {
	month, _ := strconv.Atoi(c.Query("month", "6"))
	response := controller.UserService.GetDataUser(month)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
