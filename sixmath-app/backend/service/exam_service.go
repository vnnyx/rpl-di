package service

import (
	"rpl-sixmath/entity"
	"rpl-sixmath/model"
)

type ExamService interface {
	CreateExam(request model.CreateExamRequest) (response entity.Exam, err error)
	CreateQuestion(request model.CreateQuestioRequest) (response entity.Question, err error)
}
