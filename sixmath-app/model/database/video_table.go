package database

type VideoEntity struct {
	VideoId    int             `gorm:"column:video_id;primaryKey;autoIncrement;type:int"`
	PlaylistId int             `gorm:"column:playlist_id;type:int"`
	Playlist   *PlaylistEntity `gorm:"association_foreignkey:PlaylistId"`
	Title      string          `gorm:"column:title;type:varchar(200)"`
	URLVideo   string          `gorm:"column:url_video;type:varchar(200)"`
	Deskripsi  string          `gorm:"column:deskripsi;type:longtext"`
}

func (VideoEntity) TableName() string {
	return "videos"
}
