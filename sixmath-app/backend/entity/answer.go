package entity

import "time"

type Answer struct {
	AnswerId   int       `gorm:"column:answer_id;primaryKey;autoIncrement;type:int" json:"answer_id,omitempty"`
	QuestionId int       `gorm:"column:question_id;type:int" json:"question_id,omitempty"`
	Question   *Question `gorm:"association_foreignkey:QuestionId" json:"-"`
	Answer     string    `gorm:"column:answer;type:mediumtext" json:"answer,omitempty"`
	IsTrue     bool      `gorm:"column:is_true" json:"is_true,omitempty"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime" json:"-"`
}

func (Answer) TableName() string {
	return "answers"
}
