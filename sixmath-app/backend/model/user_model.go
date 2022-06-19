package model

import (
	"mime/multipart"
	"time"
)

type StudentCreateRequest struct {
	Email                string `json:"email"`
	Username             string `json:"username"`
	Handphone            string `json:"handphone"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
	Avatar               string `json:"avatar"`
}

type StudentCreateResponse struct {
	AccessToken string    `json:"access_token"`
	UserId      int       `json:"user_id"`
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	Handphone   string    `json:"handphone"`
	Role        string    `json:"role"`
	Avatar      string    `json:"avatar"`
	CreatedAt   time.Time `json:"created_at"`
}

type TeacherCreateRequest struct {
	Email                string                `form:"email" json:"email"`
	Username             string                `form:"username" json:"username"`
	Handphone            string                `form:"handphone" json:"handphone"`
	Password             string                `form:"password" json:"password"`
	PasswordConfirmation string                `form:"password_confirmation" json:"password_confirmation"`
	Certificate          *multipart.FileHeader `form:"certificate" json:"certificate"`
	Avatar               *multipart.FileHeader `form:"avatar" json:"avatar"`
	Age                  int                   `form:"age" json:"age"`
	Description          string                `form:"description" json:"description"`
}

type TeacherCreateResponse struct {
	AccessToken string    `json:"access_token"`
	UserId      int       `json:"user_id"`
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	Handphone   string    `json:"handphone"`
	Role        string    `json:"role"`
	Certificate string    `json:"certificate"`
	Age         int       `json:"age"`
	Avatar      string    `json:"avatar"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type ParentCreateRequest struct {
	Email                string                `form:"email" json:"email"`
	Username             string                `form:"username" json:"username"`
	Handphone            string                `form:"handphone" json:"handphone"`
	Password             string                `form:"password" json:"password"`
	PasswordConfirmation string                `form:"password_confirmation" json:"password_confirmation"`
	StudentUsername      string                `form:"student_username" json:"student_username"`
	Avatar               *multipart.FileHeader `form:"avatar" json:"avatar"`
}

type ParentCreateResponse struct {
	AccessToken     string `json:"access_token"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	Handphone       string `json:"handphone"`
	Role            string `json:"role"`
	StudentUsername string `json:"student_username"`
	Avatar          string `json:"avatar"`
}

type GetUserResponse struct {
	Month  string `json:"month,omitempty"`
	Amount int64  `json:"amount"`
}

type GetTotalUser struct {
	Total int64 `json:"total,omitempty"`
}
