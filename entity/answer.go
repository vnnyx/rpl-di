package entity

type AnswerEntity struct {
	Id     uint32     `gorm:"column:answer_id"`
	ExamId uint32     `gorm:"column:exam_id"`
	Exam   ExamEntity `goem:"association_foreignKey:ExamId"`
	Answer string     `gorm:"column:answer"`
	IsTrue bool       `gorm:"column:is_true"`
}
