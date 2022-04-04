package entity

type ExamEntity struct {
	ExamId   int    `gorm:"column:exam_id;primaryKey;autoIncrement;type:int"`
	Question string `gorm:"column:question;type:mediumtext"`
}

func (ExamEntity) TableName() string {
	return "exam"
}
