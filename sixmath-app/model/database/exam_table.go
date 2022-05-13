package database

type ExamEntity struct {
	ExamId  int          `gorm:"column:exam_id;primaryKey;autoIncrement;type:int"`
	VideoId int          `gorm:"column:video_id;type:int"`
	Video   *VideoEntity `gorm:"association_foreignkey:VideoId"`
}

func (ExamEntity) TableName() string {
	return "exams"
}
