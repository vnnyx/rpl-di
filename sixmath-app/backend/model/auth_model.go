package model

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	UserID      int    `json:"user_id"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	Username    string `json:"username"`
}

type JwtPayload struct {
	UserID     int    `json:"user_id"`
	Avatar     string `json:"avatar"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	AccessUuid string `json:"access_uuid"`
}

type TokenDetails struct {
	AccessToken string
	AccessUuid  string
	AtExpires   int64
}
