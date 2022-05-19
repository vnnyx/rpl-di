package entity

type Exam struct {
	ExamId      int    `gorm:"column:exam_id;primaryKey;autoIncrement;type:int" json:"exam_id"`
	Title       string `gorm:"column:title" json:"title"`
	ImageURL    string `gorm:"column:image_url" json:"image_url"`
	Description string `gorm:"column:description" json:"description"`
	Duration    int64  `gorm:"column:duration" json:"duration"`
	VideoId     int    `gorm:"column:video_id;type:int" json:"video_id"`
	Video       *Video `gorm:"association_foreignkey:VideoId" json:"-"`
}

func (Exam) TableName() string {
	return "exams"
}
