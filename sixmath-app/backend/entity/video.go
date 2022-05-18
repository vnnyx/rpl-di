package entity

type Video struct {
	VideoId    int       `gorm:"column:video_id;primaryKey;autoIncrement;type:int" json:"video_id"`
	PlaylistId int       `gorm:"column:playlist_id;type:int" json:"playlist_id"`
	Playlist   *Playlist `gorm:"association_foreignkey:PlaylistId" json:"-"`
	Title      string    `gorm:"column:title;type:varchar(200)" json:"title"`
	URLVideo   string    `gorm:"column:url_video;type:varchar(200)" json:"url_video"`
	Deskripsi  string    `gorm:"column:deskripsi;type:longtext" json:"deskripsi"`
}

func (Video) TableName() string {
	return "videos"
}
