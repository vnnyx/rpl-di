package model

type StudentCreateRequest struct {
	Email     string `json:"email,omitempty"`
	Username  string `json:"username,omitempty"`
	Handphone string `json:"handphone,omitempty"`
	Password  string `json:"password,omitempty"`
}

type StudentCreateResponse struct {
	Email     string `json:"email,omitempty"`
	Username  string `json:"username,omitempty"`
	Handphone string `json:"handphone,omitempty"`
}
