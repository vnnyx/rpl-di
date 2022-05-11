package model

type PlaylistCreateRequest struct {
	Title     string `json:"title,omitempty"`
	TeacherId int    `json:"teacherId,omitempty"`
}

type PlaylistCreateResponse struct {
	Title     string `json:"title,omitempty"`
	TeacherId int    `json:"teacherId,omitempty"`
}
