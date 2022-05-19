package model

type CreateExamRequest struct {
	Title            string `json:"title"`
	Image            string `json:"image"`
	Description      string `json:"description"`
	DurationInMinute int64  `json:"duration_in_minute"`
	VideoId          int    `json:"video_id"`
}

type AnswerRequest struct {
	QuestionID int    `json:"question_id"`
	Answer     string `json:"answer"`
	IsTrue     bool   `json:"is_true"`
}

type CreateQuestioRequest struct {
	ExamID        int             `json:"exam_id"`
	Image         string          `json:"image"`
	Question      string          `json:"question"`
	Answers       []AnswerRequest `json:"-"`
	AnswersString string          `json:"answers"`
}
