package model

import "time"

type StudentCreateRequest struct {
	Email     string `json:"email,omitempty"`
	Username  string `json:"username,omitempty"`
	Handphone string `json:"handphone,omitempty"`
	Password  string `json:"password,omitempty"`
}

type StudentCreateResponse struct {
	UserId    int       `json:"user_id,omitempty"`
	Email     string    `json:"email,omitempty"`
	Username  string    `json:"username,omitempty"`
	Handphone string    `json:"handphone,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type GetUserResponse struct {
	Month  string `json:"month,omitempty"`
	Amount int64  `json:"amount"`
}

type GetTotalUser struct {
	Total int64 `json:"total,omitempty"`
}
