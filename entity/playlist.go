package entity

type PlaylistEntity struct {
	PlaylistId int            `gorm:"column:playlist_id;primaryKey;autoIncrement;type:int"`
	Title      string         `gorm:"column:title;type:mediumtext"`
	TeacherId  int            `gorm:"column:teacher_id;type:int"`
	Teacher    *TeacherEntity `gorm:"association_foreignKey:TeacherId"`
}

func (PlaylistEntity) TableName() string {
	return "playlist"
}
