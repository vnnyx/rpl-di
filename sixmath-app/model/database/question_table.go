package database

type QuestionEntity struct {
	QuestionId int         `gorm:"column:question_id;primaryKey;autoIncrement;type:int"`
	Question   string      `gorm:"column:question;type:varchar(100)"`
	ExamId     int         `gorm:"column:exam_id;type:int"`
	Exam       *ExamEntity `gorm:"association_foreignkey:ExamId"`
}

func (QuestionEntity) TableName() string {
	return "questions"
}
