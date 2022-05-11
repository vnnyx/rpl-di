package controller

import (
	"rpl-sixmath/exception"
	"rpl-sixmath/model"
	"rpl-sixmath/service"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(service *service.UserService) UserController {
	return UserController{UserService: *service}
}

func (controller *UserController) Route(app *fiber.App) {
	app.Post("/api/student", controller.CreateStudent)
}

func (controller *UserController) CreateStudent(c *fiber.Ctx) error {
	var request model.StudentCreateRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)
	response := controller.UserService.CreateStudent(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
