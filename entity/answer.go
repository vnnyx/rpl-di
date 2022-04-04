package entity

type AnswerEntity struct {
	Id     int
	ExamId ExamEntity
	Answer string
	IsTrue bool
}
