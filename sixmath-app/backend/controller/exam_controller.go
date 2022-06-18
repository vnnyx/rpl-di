package controller

import (
	"rpl-sixmath/exception"
	"rpl-sixmath/helper"
	"rpl-sixmath/middleware"
	"rpl-sixmath/model"
	"rpl-sixmath/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ExamController struct {
	service service.ExamService
}

func NewExamController(service service.ExamService) ExamController {
	return ExamController{service: service}
}

func (controller ExamController) Route(app fiber.Router) {
	router := app.Group("/api/exam", middleware.CheckToken())

	router.Post("/", middleware.IsTeacher(), controller.CreateExam)
	router.Post("/question/:exam_id", middleware.IsTeacher(), controller.CreateQuestion)
}

func (controller ExamController) CreateExam(c *fiber.Ctx) error {
	title := c.FormValue("title")
	desc := c.FormValue("description")
	duration, _ := strconv.Atoi(c.FormValue("duration_in_minute"))

	image, err := c.FormFile("image")
	exception.PanicIfNeeded(err)
	result, err := helper.UploadToCloudinary(c, image, ".:/uploads/exam/", image.Filename)
	exception.PanicIfNeeded(err)

	imageUrl := result.SecureURL

	response, err := controller.service.CreateExam(model.CreateExamRequest{
		Title:            title,
		Image:            imageUrl,
		Description:      desc,
		DurationInMinute: int64(duration),
	})
	exception.PanicIfNeeded(err)

	return c.Status(200).JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})

}

func (controller ExamController) CreateQuestion(c *fiber.Ctx) error {
	question := c.FormValue("question")
	answers := c.FormValue("answers")
	examId, _ := strconv.Atoi(c.Params("exam_id"))

	image, err := c.FormFile("image")
	exception.PanicIfNeeded(err)
	result, err := helper.UploadToCloudinary(c, image, ".:/sixmath/exam/", image.Filename)
	exception.PanicIfNeeded(err)
	imageUrl := result.SecureURL

	response, err := controller.service.CreateQuestion(model.CreateQuestionRequest{
		ExamID:        examId,
		Image:         imageUrl,
		Question:      question,
		AnswersString: answers,
	})
	exception.PanicIfNeeded(err)

	return c.Status(200).JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})

}
