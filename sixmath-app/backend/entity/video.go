package entity

import "time"

type Video struct {
	VideoId         int       `gorm:"column:video_id;primaryKey;autoIncrement;type:int" json:"video_id"`
	TeacherUsername string    `gorm:"column:teacher;type:varchar(100)" json:"teacher"`
	Teacher         *User     `gorm:"association_foreignkey:TeacherUsername;references:Username" json:"-"`
	Title           string    `gorm:"column:title;type:varchar(200)" json:"title"`
	URLVideo        string    `gorm:"column:url_video;type:varchar(200)" json:"url_video"`
	Description     string    `gorm:"column:description;type:longtext" json:"description"`
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime" json:"-"`
}

func (Video) TableName() string {
	return "videos"
}
