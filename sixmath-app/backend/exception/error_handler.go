package exception

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"rpl-sixmath/model"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(ValidationError)
	if ok {
		var obj interface{}
		_ = json.Unmarshal([]byte(err.Error()), &obj)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   nil,
			Error:  obj,
		})
	}
	if err.Error() == "USERNAME_REGISTERED" {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.Response{
			Code:   fiber.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   nil,
			Error: map[string]interface{}{
				"username": "MUST_UNIQUE",
			},
		})
	}

	if err.Error() == "PASSWORD_NOT_MATCH" {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.Response{
			Code:   fiber.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   nil,
			Error: map[string]interface{}{
				"password": "NOT_MATCH",
			},
		})
	}

	if err.Error() == "PLAYLIST_NOT_FOUND" {
		return ctx.Status(fiber.StatusNotFound).JSON(model.Response{
			Code:   fiber.StatusNotFound,
			Status: "BAD_REQUEST",
			Data:   nil,
			Error: map[string]interface{}{
				"playlist_id": "NOT_FOUND",
			},
		})
	}

	if err.Error() == "VIDEO_NOT_FOUND" {
		return ctx.Status(fiber.StatusNotFound).JSON(model.Response{
			Code:   fiber.StatusNotFound,
			Status: "BAD_REQUEST",
			Data:   nil,
			Error: map[string]interface{}{
				"video_id": "NOT_FOUND",
			},
		})
	}

	if err.Error() == "STUDENT_NOT_FOUND" {
		return ctx.Status(fiber.StatusNotFound).JSON(model.Response{
			Code:   fiber.StatusNotFound,
			Status: "BAD_REQUEST",
			Data:   nil,
			Error: map[string]interface{}{
				"student_username": "NOT_FOUND",
			},
		})
	}

	if err.Error() == "USER_NOT_FOUND" {
		return ctx.Status(fiber.StatusNotFound).JSON(model.Response{
			Code:   fiber.StatusNotFound,
			Status: "BAD_REQUEST",
			Data:   nil,
			Error: map[string]interface{}{
				"username": "NOT_FOUND",
			},
		})
	}

	if err.Error() == "EMAIL_NOT_FOUND" {
		return ctx.Status(fiber.StatusNotFound).JSON(model.Response{
			Code:   fiber.StatusNotFound,
			Status: "BAD_REQUEST",
			Data:   nil,
			Error: map[string]interface{}{
				"email": "NOT_FOUND",
			},
		})
	}

	if err.Error() == "TOKEN_INVALID" {
		return ctx.Status(fiber.StatusForbidden).JSON(model.Response{
			Code:   fiber.StatusForbidden,
			Status: "FORBIDDEN",
			Data:   nil,
			Error: map[string]interface{}{
				"otp": "INVALID",
			},
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse{
		Code:   500,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err.Error(),
	})
}
