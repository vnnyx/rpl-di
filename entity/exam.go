package entity

type ExamEntity struct {
	Id       int    `gorm:"column:exam_id;primary;autoIncrement"`
	Question string `gorm:"column:question"`
}

func (ExamEntity) TableName() string {
	return "exam"
}
