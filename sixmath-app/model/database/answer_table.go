package database

type AnswerEntity struct {
	AnswerId   int             `gorm:"column:answer_id;primaryKey;autoIncrement;type:int"`
	QuestionId int             `gorm:"column:question_id;type:int"`
	Question   *QuestionEntity `gorm:"association_foreignkey:QuestionId"`
	Answer     string          `gorm:"column:answer;type:mediumtext"`
	IsTrue     bool            `gorm:"column:is_true"`
}

func (AnswerEntity) TableName() string {
	return "answers"
}
