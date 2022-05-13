package entity

type User struct {
	UserId          int    `json:"user_id,omitempty"`
	Username        string `json:"username,omitempty"`
	Email           string `json:"email,omitempty"`
	Handphone       string `json:"handphone,omitempty"`
	Password        string `json:"password,omitempty"`
	Role            string `json:"role,omitempty"`
	Gender          string `json:"gender,omitempty"`
	Certificate     string `json:"certificate,omitempty"`
	Avatar          string `json:"avatar,omitempty"`
	StudentUsername string `json:"student_username,omitempty"`
}
