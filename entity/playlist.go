package entity

type PlaylistEntity struct {
	Id        int           `gorm:"column:playlist_id;primaryKey;autoIncrement"`
	Title     string        `gorm:"column:title"`
	TeacherId int           `gorm:"column:teacher_id"`
	Teacher   TeacherEntity `gorm:"association_foreignKey:TeacherId"`
}

func (PlaylistEntity) TableName() string {
	return "playlist"
}
