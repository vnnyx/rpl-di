package entity

type ExamEntity struct {
	ExamId   int    `gorm:"column:exam_id;primaryKey;autoIncrement"`
	Question string `gorm:"column:question"`
}

func (ExamEntity) TableName() string {
	return "exam"
}
