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
			Data:   struct{}{},
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

	return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse{
		Code:   500,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err.Error(),
	})
}
