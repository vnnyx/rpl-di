package service

import (
	"rpl-sixmath/entity"
	"rpl-sixmath/model"
)

type ExamService interface {
	CreateExam(request model.CreateExamRequest) (response entity.Exam, err error)
	CreateQuestion(request model.CreateQuestionRequest) (response entity.Question, err error)
	GetAllExam(orderBy string) (response []model.GetExamResponse, err error)
}
