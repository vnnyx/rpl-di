package repository

import (
	"rpl-sixmath/entity"
	"rpl-sixmath/model"

	"gorm.io/gorm"
)

type examRepositoryImpl struct {
	DB *gorm.DB
}

func NewExamRepository(DB *gorm.DB) ExamRepository {
	return &examRepositoryImpl{DB: DB}
}

func (repo examRepositoryImpl) InsertExam(request entity.Exam) (entity.Exam, error) {
	err := repo.DB.Create(&request).Error
	return request, err
}

func (repo examRepositoryImpl) InsertQuestion(request model.CreateQuestionRequest) (entity.Question, error) {
	tx := repo.DB.Begin()

	question := entity.Question{
		ExamId:   request.ExamID,
		Question: request.Question,
		Image:    request.Image,
	}

	err := tx.Create(&question).Error

	if err != nil {
		tx.Rollback()
		return entity.Question{}, err
	}
	for _, v := range request.Answers {
		err = tx.Create(&entity.Answer{
			QuestionId: question.QuestionId,
			Answer:     v.Answer,
			IsTrue:     v.IsTrue,
		}).Error

		if err != nil {
			tx.Rollback()
			return entity.Question{}, err
		}
	}

	err = tx.Commit().Error

	return question, err

}

func (repo examRepositoryImpl) GetExamByID(examID int) (response entity.Exam, err error) {
	err = repo.DB.Where("exam_id", examID).First(&response).Error

	return response, err
}

func (repo examRepositoryImpl) GetAllExam(orderBy string) (response []entity.Exam, err error) {
	if orderBy == "alphabet" {
		err = repo.DB.Order("title").Find(&response).Error
	} else {
		err = repo.DB.Order("created_at desc").Find(&response).Error
	}
	return response, err
}

func (repo examRepositoryImpl) GetTotalQuestion(examId int) (response int64, err error) {
	err = repo.DB.Find(&entity.Question{}).Where("exam_id", examId).Count(&response).Error
	return response, err
}
