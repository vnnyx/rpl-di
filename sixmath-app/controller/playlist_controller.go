package controller

import (
	"fmt"
	"rpl-sixmath/exception"
	"rpl-sixmath/middleware"
	"rpl-sixmath/model"
	"rpl-sixmath/service"

	"github.com/gofiber/fiber/v2"
)

type PlaylistController struct {
	PlaylistService service.PlaylistService
}

func NewPlaylistController(playlistService *service.PlaylistService) PlaylistController {
	return PlaylistController{PlaylistService: *playlistService}
}

func (controller *PlaylistController) Route(app *fiber.App) {
	router := app.Group("/api/playlist", middleware.CheckToken())
	router.Post("/create", controller.CreatePlaylist)
}

func (controller *PlaylistController) CreatePlaylist(c *fiber.Ctx) error {
	userId := c.Locals("currentUserId").(int)
	fmt.Print(userId)
	var request model.PlaylistCreateRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)
	response := controller.PlaylistService.CreatePlaylist(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
