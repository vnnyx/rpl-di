package exception

import (
	"github.com/gofiber/fiber/v2"
	"rpl-sixmath/model"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(ValidationError)
	if ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse{
		Code:   500,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err.Error(),
	})
}
