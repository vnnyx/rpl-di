package repository

import (
	"rpl-sixmath/entity"
	"rpl-sixmath/model"
)

type ExamRepository interface {
	GetExamByID(examID int) (response entity.Exam, err error)
	InsertExam(request entity.Exam) (entity.Exam, error)
	InsertQuestion(request model.CreateQuestioRequest) (entity.Question, error)
}
