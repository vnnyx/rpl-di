package entity

type Playlist struct {
	PlaylistId int    `json:"playlist_id,omitempty"`
	Title      string `json:"title,omitempty"`
	TeacherId  int    `json:"teacher_id,omitempty"`
}
