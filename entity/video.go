package entity

type VideoEntity struct {
	VideoId    int             `gorm:"column:video_id;primaryKey;autoIncrement;type:int"`
	Playlist   *PlaylistEntity `gorm:"association_foreignKey:PlaylistId"`
	PlaylistId int             `gorm:"column:playlist_id;type:int"`
	Exam       *ExamEntity     `gorm:"association_foreignKey:ExamId"`
	ExamId     int             `gorm:"column:exam_id;type:int"`
	Title      string          `gorm:"column:title;type:varchar(200)"`
	URLVideo   string          `gorm:"column:url_video;type:varchar(200)"`
	Deskripsi  string          `gorm:"column:deskripsi;type:longtext"`
}

func (VideoEntity) TableName() string {
	return "video"
}
