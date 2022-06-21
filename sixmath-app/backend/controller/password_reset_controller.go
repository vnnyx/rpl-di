package controller

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"rpl-sixmath/exception"
	"rpl-sixmath/model"
	"rpl-sixmath/service"
	"strconv"
)

type PasswordResetController struct {
	service.PasswordResetService
}

func NewPasswordResetController(passwordResetService *service.PasswordResetService) PasswordResetController {
	return PasswordResetController{PasswordResetService: *passwordResetService}
}

func (controller *PasswordResetController) Route(app *fiber.App) {
	router := app.Group("/api/password-reset")
	router.Post("/send-otp", controller.SendOtp)
	router.Post("/validate-otp", controller.ValidateOtp)
	router.Post("/", controller.PasswordReset)
}

func (controller *PasswordResetController) SendOtp(ctx *fiber.Ctx) error {
	var request model.SendOtpRequest
	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)
	response, err := controller.PasswordResetService.SendOtp(request)
	exception.PanicIfNeeded(err)
	return ctx.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *PasswordResetController) ValidateOtp(ctx *fiber.Ctx) error {
	var request model.SendOtpRequest
	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)
	email := ctx.Query("email")
	request.Email = email
	err = controller.PasswordResetService.ValidateOtp(request)
	exception.PanicIfNeeded(err)
	return ctx.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
	})
}

func (controller *PasswordResetController) PasswordReset(ctx *fiber.Ctx) error {
	var request model.PasswordResetRequest
	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)
	if request.Password != request.PasswordConfirmation {
		return errors.New("PASSWORD_NOT_MATCH")
	}
	userId, _ := strconv.Atoi(ctx.Query("id"))
	request.UserId = userId
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	request.Password = string(passwordHash)
	err = controller.PasswordResetService.PasswordReset(request)
	exception.PanicIfNeeded(err)
	return ctx.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
	})
}
