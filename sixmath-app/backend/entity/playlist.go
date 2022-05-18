package entity

type Playlist struct {
	PlaylistId int    `gorm:"column:playlist_id;primaryKey;autoIncrement;type:int" json:"playlist_id,omitempty"`
	Title      string `gorm:"column:title;type:mediumtext" json:"title,omitempty"`
	TeacherId  int    `gorm:"column:teacher_id;type:int" json:"teacher_id,omitempty"`
	Teacher    *User  `gorm:"association_foreignkey:UserId" json:"-"`
}

func (Playlist) TableName() string {
	return "playlists"
}
