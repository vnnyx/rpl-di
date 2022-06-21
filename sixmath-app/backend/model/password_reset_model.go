package model

type SendOtpRequest struct {
	Email string `json:"email"`
	Otp   int    `json:"otp"`
}

type PasswordResetRequest struct {
	UserId               int    `json:"-"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type PasswordResetResponse struct {
	Email string `json:"email"`
}
