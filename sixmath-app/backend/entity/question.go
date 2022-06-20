package entity

import "time"

type Question struct {
	QuestionId int       `gorm:"column:question_id;primaryKey;autoIncrement;type:int" json:"question_id,omitempty"`
	Question   string    `gorm:"column:question;type:varchar(100)" json:"question,omitempty"`
	Image      string    `gorm:"column:image;type:varchar(255)" json:"image"`
	ExamId     int       `gorm:"column:exam_id;type:int" json:"exam_id,omitempty"`
	Exam       *Exam     `gorm:"association_foreignkey:ExamId" json:"-"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime" json:"-"`
}

func (Question) TableName() string {
	return "questions"
}
