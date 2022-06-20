package entity

import "time"

type Exam struct {
	ExamId          int       `gorm:"column:exam_id;primaryKey;autoIncrement;type:int" json:"exam_id"`
	TeacherUsername string    `gorm:"column:teacher;type:varchar(100)" json:"teacher"`
	Teacher         *User     `gorm:"association_foreignkey:TeacherUsername;references:Username" json:"-"`
	Title           string    `gorm:"column:title" json:"title"`
	ImageURL        string    `gorm:"column:image_url" json:"image_url"`
	Description     string    `gorm:"column:description" json:"description"`
	Duration        int64     `gorm:"column:duration" json:"duration"`
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime" json:"-"`
}

func (Exam) TableName() string {
	return "exams"
}
