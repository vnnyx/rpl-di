package entity

type Exam struct {
	ExamId  int    `gorm:"column:exam_id;primaryKey;autoIncrement;type:int" json:"exam_id,omitempty"`
	VideoId int    `gorm:"column:video_id;type:int" json:"video_id,omitempty"`
	Video   *Video `gorm:"association_foreignkey:VideoId" json:"-"`
}

func (Exam) TableName() string {
	return "exams"
}
