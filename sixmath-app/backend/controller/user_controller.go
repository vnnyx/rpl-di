package controller

import (
	"rpl-sixmath/exception"
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
	router.Post("/parent", controller.CreateParent)
	router.Get("/teacher/all", middleware.CheckToken(), controller.GetAllTeacher)
	router.Get("/registration-statistics", middleware.CheckToken(), controller.GetDataUser)
}

func (controller *UserController) CreateStudent(c *fiber.Ctx) error {
	var request model.StudentCreateRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	if request.Password != request.PasswordConfirmation {
		exception.PanicIfNeeded("PASSWORD_NOT_MATCH")
	}
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

	request.Certificate = certificate
	request.Avatar = avatar

	if request.Password != request.PasswordConfirmation {
		exception.PanicIfNeeded("PASSWORD_NOT_MATCH")
	}
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

func (controller *UserController) CreateParent(c *fiber.Ctx) error {
	var request model.ParentCreateRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	avatar, _ := c.FormFile("avatar")
	request.Avatar = avatar

	if request.Password != request.PasswordConfirmation {
		exception.PanicIfNeeded("PASSWORD_NOT_MATCH")
	}
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	request.Password = string(passwordHash)

	response, err := controller.UserService.CreateParent(request)
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
