package controller

import (
	"rpl-sixmath/exception"
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
	router := app.Group("/api/exam")

	router.Post("/", controller.CreateExam)
	router.Post("/question/:exam_id", controller.CreateQuestion)
}

func (controller ExamController) CreateExam(c *fiber.Ctx) error {
	title := c.FormValue("title")
	desc := c.FormValue("description")
	duration, _ := strconv.Atoi(c.FormValue("duration_in_minute"))
	videoID, _ := strconv.Atoi(c.FormValue("video_id"))

	image, _ := c.FormFile("image")

	response, err := controller.service.CreateExam(model.CreateExamRequest{
		Title:            title,
		Image:            "/uploads/image/" + image.Filename,
		Description:      desc,
		DurationInMinute: int64(duration),
		VideoId:          videoID,
	})
	exception.PanicIfNeeded(err)

	err = c.SaveFile(image, "/uploads/image/"+image.Filename)
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

	response, err := controller.service.CreateQuestion(model.CreateQuestioRequest{
		ExamID:        examId,
		Image:         "/uploads/image/" + image.Filename,
		Question:      question,
		AnswersString: answers,
	})
	exception.PanicIfNeeded(err)

	c.SaveFile(image, "./uploads/image/"+image.Filename)
	return c.Status(200).JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})

}
