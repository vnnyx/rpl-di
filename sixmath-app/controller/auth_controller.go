package controller

import (
	"rpl-sixmath/exception"
	"rpl-sixmath/model"
	"rpl-sixmath/service"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(authService *service.AuthService) AuthController {
	return AuthController{AuthService: *authService}
}

func (controller AuthController) Route(app fiber.Router) {
	router := app.Group("/api/auth")

	router.Post("/login", controller.Login)

}

func (controller AuthController) Login(c *fiber.Ctx) error {
	request := new(model.LoginRequest)

	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response, err := controller.AuthService.Login(*request)
	exception.PanicIfNeeded(err)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})

}
