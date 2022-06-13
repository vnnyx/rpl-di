package model

import "time"

type StudentCreateRequest struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Handphone string `json:"handphone"`
	Avatar    string `json:"avatar"`
	Password  string `json:"password"`
}

type StudentResponse struct {
	UserId    int       `json:"user_id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Handphone string    `json:"handphone"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
}

type TeacherCreateRequest struct {
	Email       string `json:"email"`
	Username    string `json:"username"`
	Handphone   string `json:"handphone"`
	Password    string `json:"password"`
	Certificate string `json:"certificate"`
	Avatar      string `json:"avatar"`
}

type TeacherResponse struct {
	UserId      int       `json:"user_id"`
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	Handphone   string    `json:"handphone"`
	Certificate string    `json:"certificate"`
	Avatar      string    `json:"avatar"`
	CreatedAt   time.Time `json:"created_at"`
}

type GetUserResponse struct {
	Month  string `json:"month,omitempty"`
	Amount int64  `json:"amount"`
}

type GetTotalUser struct {
	Total int64 `json:"total,omitempty"`
}
