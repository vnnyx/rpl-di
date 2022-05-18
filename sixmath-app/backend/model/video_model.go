package model

type VideoCreateRequest struct {
	PlaylistId int    `json:"playlist_id,omitempty"`
	Title      string `json:"title,omitempty"`
	URLVideo   string `json:"url_video,omitempty"`
	Deskripsi  string `json:"deskripsi,omitempty"`
}

type VideoUpdateRequest struct {
	VideoId    int    `json:"video_id,omitempty"`
	PlaylistId int    `json:"playlist_id,omitempty"`
	Title      string `json:"title,omitempty"`
	URLVideo   string `json:"url_video,omitempty"`
	Deskripsi  string `json:"deskripsi,omitempty"`
}

type VideoResponse struct {
	VideoId    int    `json:"video_id,omitempty"`
	PlaylistId int    `json:"playlist_id,omitempty"`
	Title      string `json:"title,omitempty"`
	URLVideo   string `json:"url_video,omitempty"`
	Deskripsi  string `json:"deskripsi,omitempty"`
}
