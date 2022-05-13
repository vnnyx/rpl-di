package entity

type Question struct {
	QuestionId int    `json:"question_id,omitempty"`
	Question   string `json:"question,omitempty"`
	ExamId     int    `json:"exam_id,omitempty"`
}
