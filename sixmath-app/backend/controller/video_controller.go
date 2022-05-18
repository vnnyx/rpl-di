package controller

import (
	"github.com/gofiber/fiber/v2"
	"rpl-sixmath/exception"
	"rpl-sixmath/middleware"
	"rpl-sixmath/model"
	"rpl-sixmath/service"
	"strconv"
)

type VideoController struct {
	VideoService service.VideoService
}

func NewVideoController(videoService *service.VideoService) VideoController {
	return VideoController{VideoService: *videoService}
}

func (controller *VideoController) Route(app *fiber.App) {
	router := app.Group("/api/video", middleware.CheckToken())
	router.Post("/", controller.CreateVideo)
	router.Get("/", controller.MainVideo)
	router.Get("/recommended", controller.RecommendedVideo)
	router.Delete("/:id", controller.DeleteVideo)
	router.Put("/:id", controller.UpdateVideo)
}

func (controller *VideoController) CreateVideo(c *fiber.Ctx) error {
	var request model.VideoCreateRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response := controller.VideoService.CreateVideo(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *VideoController) MainVideo(c *fiber.Ctx) error {
	response := controller.VideoService.GetMainVideo()
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller VideoController) RecommendedVideo(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	var pagination model.Pagination
	pagination.Page = page
	pagination.Limit = 3
	response := controller.VideoService.GetRecommendedVideo(pagination)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller VideoController) DeleteVideo(c *fiber.Ctx) error {
	videoId := c.Params("id")
	id, _ := strconv.Atoi(videoId)
	controller.VideoService.DeleteVideo(id)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
	})
}

func (controller VideoController) UpdateVideo(c *fiber.Ctx) error {
	videoId := c.Params("id")
	id, _ := strconv.Atoi(videoId)
	var request model.VideoUpdateRequest
	request.VideoId = id
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response := controller.VideoService.UpdateVideo(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
