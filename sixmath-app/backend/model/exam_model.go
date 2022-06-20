package model

type CreateExamRequest struct {
	Teacher          string `json:"teacher"`
	Title            string `json:"title"`
	Image            string `json:"image"`
	Description      string `json:"description"`
	DurationInMinute int64  `json:"duration_in_minute"`
}

type GetExamResponse struct {
	ExamId           int    `json:"exam_id"`
	Teacher          string `json:"teacher"`
	Title            string `json:"title"`
	Image            string `json:"image"`
	Description      string `json:"description"`
	DurationInMinute int64  `json:"duration_in_minute"`
	Total            int64  `json:"total_question"`
}

type AnswerRequest struct {
	QuestionID int    `json:"question_id"`
	Answer     string `json:"answer"`
	IsTrue     bool   `json:"is_true"`
}

type CreateQuestionRequest struct {
	ExamID        int             `json:"exam_id"`
	Image         string          `json:"image"`
	Question      string          `json:"question"`
	Answers       []AnswerRequest `json:"-"`
	AnswersString string          `json:"answers"`
}
