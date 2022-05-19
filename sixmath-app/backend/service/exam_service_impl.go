package service

import (
	"encoding/json"
	"rpl-sixmath/entity"
	"rpl-sixmath/model"
	"rpl-sixmath/repository"
	"rpl-sixmath/validation"
	"time"
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

	//check video
	_, err = service.videoRepo.FindVideoById(request.VideoId)
	if err != nil {
		return response, err
	}

	response, err = service.examRepo.InsertExam(entity.Exam{
		Title:       request.Title,
		ImageURL:    request.Image,
		Description: request.Description,
		Duration:    request.DurationInMinute * int64(time.Minute),
		VideoId:     request.VideoId,
	})

	return response, err
}

func (service examServiceImpl) CreateQuestion(request model.CreateQuestioRequest) (response entity.Question, err error) {
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
