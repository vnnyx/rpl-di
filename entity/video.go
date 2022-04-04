package entity

type VideoEntity struct {
	VideoId    int             `gorm:"column:video_id;primaryKey;autoIncrement"`
	Playlist   *PlaylistEntity `gorm:"association_foreignKey:PlaylistId"`
	PlaylistId int             `gorm:"column:playlist_id"`
	Exam       *ExamEntity     `gorm:"association_foreignKey:ExamId"`
	ExamId     int             `gorm:"column:exam_id"`
	Title      string          `gorm:"column:title"`
	URLVideo   string          `gorm:"column:url_video"`
	Deskripsi  string          `gorm:"column:deskripsi"`
}

func (VideoEntity) TableName() string {
	return "video"
}
