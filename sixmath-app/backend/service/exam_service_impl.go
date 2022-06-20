package service

import (
	"encoding/json"
	"rpl-sixmath/entity"
	"rpl-sixmath/exception"
	"rpl-sixmath/model"
	"rpl-sixmath/repository"
	"rpl-sixmath/validation"
)

type examServiceImpl struct {
	examRepo  repository.ExamRepository
	videoRepo repository.VideoRepository
}

func NewExamService(examRepo repository.ExamRepository, videoRepo repository.VideoRepository) ExamService {
	return &examServiceImpl{examRepo: examRepo, videoRepo: videoRepo}
}

func (service examServiceImpl) CreateExam(request model.CreateExamRequest) (response entity.Exam, err error) {
	//validation
	validation.CreateExamValidate(request)

	if err != nil {
		return response, err
	}

	response, err = service.examRepo.InsertExam(entity.Exam{
		TeacherUsername: request.Teacher,
		Title:           request.Title,
		ImageURL:        request.Image,
		Description:     request.Description,
		Duration:        request.DurationInMinute,
	})

	return response, err
}

func (service examServiceImpl) CreateQuestion(request model.CreateQuestionRequest) (response entity.Question, err error) {
	//validaiton

	//check exam
	_, err = service.examRepo.GetExamByID(request.ExamID)
	if err != nil {
		return response, err
	}

	answers := make([]model.AnswerRequest, 0)
	json.Unmarshal([]byte(request.AnswersString), &answers)
	request.Answers = answers
	response, err = service.examRepo.InsertQuestion(request)

	return response, err

}

func (service examServiceImpl) GetAllExam(orderBy string) (response []model.GetExamResponse, err error) {
	exams, err := service.examRepo.GetAllExam(orderBy)
	if err != nil {
		return response, err
	}
	for _, exam := range exams {
		count, err := service.examRepo.GetTotalQuestion(exam.ExamId)
		exception.PanicIfNeeded(err)
		response = append(response, model.GetExamResponse{
			ExamId:           exam.ExamId,
			Teacher:          exam.TeacherUsername,
			Title:            exam.Title,
			Image:            exam.ImageURL,
			Description:      exam.Description,
			DurationInMinute: exam.Duration,
			Total:            count,
		})
	}

	return response, nil
}
