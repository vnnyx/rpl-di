package repository

import (
	"rpl-sixmath/entity"
	"rpl-sixmath/model"
)

type ExamRepository interface {
	GetExamByID(examID int) (response entity.Exam, err error)
	GetAllExam(orderBy string) (response []entity.Exam, err error)
	InsertExam(request entity.Exam) (entity.Exam, error)
	GetTotalQuestion(examId int) (response int64, err error)
	InsertQuestion(request model.CreateQuestionRequest) (entity.Question, error)
}
