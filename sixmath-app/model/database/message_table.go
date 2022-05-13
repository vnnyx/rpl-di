package database

import (
	"time"
)

type MessageEntity struct {
	MessageId int         `gorm:"column:message_id;primaryKey;type:int;autoIncrement"`
	Message   string      `gorm:"column:message;type:longtext"`
	TargetId  string      `gorm:"column:target_id;type:int"`
	Time      time.Time   `gorm:"column:time"`
	Action    string      `gorm:"column:action;type:varchar(100)"`
	StudentId int         `gorm:"column:client_id;type:int"`
	Student   *UserEntity `gorm:"association_foreignkey:UserId"`
}

func (MessageEntity) TableName() string {
	return "messages"
}
