package entity

import "time"

type MessageEntity struct {
	MessageId int                `gorm:"column:message_id;primaryKey;autoIncrement"`
	Message   string             `gorm:"column:message"`
	TargetId  int                `gorm:"column:target_id"`
	Target    *CodeMessageEntity `gorm:"association_foreignKey:Id"`
	TeacherId int                `gorm:"column:teacher_id"`
	Teacher   *TeacherEntity     `gorm:"association_foreignKey:TeacherId"`
	StudentId int                `gorm:"column:student_id"`
	Student   *StudentEntity     `gorm:"association_foreignKey:StudentId"`
	File      string             `gorm:"column:file"`
	Time      time.Time          `gorm:"column:time"`
	Action    string             `gorm:"column:action"`
}

func (MessageEntity) TableName() string {
	return "message"
}
