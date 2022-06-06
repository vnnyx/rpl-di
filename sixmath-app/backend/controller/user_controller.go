package controller

import (
	"rpl-sixmath/exception"
	"rpl-sixmath/middleware"
	"rpl-sixmath/model"
	"rpl-sixmath/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
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
	app.Get("/api/dashboard/statistik-pendaftaran", middleware.CheckToken(), controller.GetDataUser)
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

func (controller *UserController) GetDataUser(c *fiber.Ctx) error {
	month, _ := strconv.Atoi(c.Query("month", "6"))
	response := controller.UserService.GetDataUser(month)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
