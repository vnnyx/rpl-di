package model

type PlaylistCreateRequest struct {
	Title     string `json:"title,omitempty"`
	TeacherId int    `json:"teacherId,omitempty"`
}

type PlaylistCreateResponse struct {
	PlaylistId int    `json:"playlist_id,omitempty"`
	Title      string `json:"title,omitempty"`
	TeacherId  int    `json:"teacherId,omitempty"`
}
