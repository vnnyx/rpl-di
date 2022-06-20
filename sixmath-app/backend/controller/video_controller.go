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
	router.Post("/create", middleware.IsTeacher(), controller.CreateVideo)
	router.Get("/detail/:id", controller.DetailVideo)
	router.Get("/recommended", controller.RecommendedVideo)
	router.Delete("/:id", middleware.IsTeacher(), controller.DeleteVideo)
	router.Put("/:id", middleware.IsTeacher(), controller.UpdateVideo)
	router.Get("/all", controller.GetAllVideo)
}

func (controller *VideoController) CreateVideo(c *fiber.Ctx) error {
	teacher := c.Locals("currentUsername").(string)
	var request model.VideoCreateRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	request.Teacher = teacher
	response, err := controller.VideoService.CreateVideo(request)
	exception.PanicIfNeeded(err)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *VideoController) DetailVideo(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	response, err := controller.VideoService.DetailVideo(id)
	exception.PanicIfNeeded(err)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller VideoController) RecommendedVideo(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	main, _ := strconv.Atoi(c.Query("main"))
	var pagination model.Pagination
	pagination.Page = page
	pagination.MainVideoId = main
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
	teacher := c.Locals("currentUsername").(string)
	id, _ := strconv.Atoi(videoId)
	var request model.VideoUpdateRequest
	request.VideoId = id
	request.Teacher = teacher
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response, err := controller.VideoService.UpdateVideo(request)
	exception.PanicIfNeeded(err)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller VideoController) GetAllVideo(c *fiber.Ctx) error {
	orderBy := c.Query("orderBy", "new")
	response, err := controller.VideoService.GetAllVideo(orderBy)
	exception.PanicIfNeeded(err)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
