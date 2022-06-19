package model

type VideoCreateRequest struct {
	PlaylistId  int    `json:"playlist_id"`
	Teacher     string `json:"teacher"`
	Title       string `json:"title"`
	URLVideo    string `json:"url_video"`
	Description string `json:"description"`
}

type VideoUpdateRequest struct {
	VideoId     int    `json:"video_id"`
	Teacher     string `json:"teacher"`
	Title       string `json:"title"`
	URLVideo    string `json:"url_video"`
	Description string `json:"description"`
}

type VideoResponse struct {
	VideoId     int    `json:"video_id"`
	Teacher     string `json:"teacher"`
	Title       string `json:"title"`
	URLVideo    string `json:"url_video"`
	Description string `json:"description"`
}
