package entity

type Answer struct {
	AnswerId   int    `json:"answer_id,omitempty"`
	QuestionId int    `json:"question_id,omitempty"`
	Answer     string `json:"answer,omitempty"`
	IsTrue     bool   `json:"is_true,omitempty"`
}
