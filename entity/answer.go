package entity

type AnswerEntity struct {
	AnswerId int         `gorm:"column:answer_id;primaryKey;autoIncrement;type:int"`
	ExamId   int         `gorm:"column:exam_id;type:int"`
	Exam     *ExamEntity `gorm:"association_foreignKey:ExamId"`
	Answer   string      `gorm:"column:answer;type:mediumtext"`
	IsTrue   bool        `gorm:"column:is_true;type:bool"`
}

func (AnswerEntity) TableName() string {
	return "answer"
}
