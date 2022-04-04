package entity

type AnswerEntity struct {
	AnswerId int         `gorm:"column:answer_id;primaryKey;autoIncrement"`
	ExamId   int         `gorm:"column:exam_id"`
	Exam     *ExamEntity `gorm:"association_foreignKey:ExamId"`
	Answer   string      `gorm:"column:answer"`
	IsTrue   bool        `gorm:"column:is_true"`
}

func (AnswerEntity) TableName() string {
	return "answer"
}
