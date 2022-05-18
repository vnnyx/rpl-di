package entity

import (
	"time"
)

type Message struct {
	MessageId int       `gorm:"column:message_id;primaryKey;type:int;autoIncrement" json:"message_id,omitempty"`
	Message   string    `gorm:"column:message;type:longtext" json:"message,omitempty"`
	TargetId  string    `gorm:"column:target_id;type:int" json:"target_id,omitempty"`
	Time      time.Time `gorm:"column:time" json:"time"`
	Action    string    `gorm:"column:action;type:varchar(100)" json:"action,omitempty"`
	StudentId int       `gorm:"column:client_id;type:int" json:"student_id,omitempty"`
	Student   *User     `gorm:"association_foreignkey:UserId" json:"-"`
}

func (Message) TableName() string {
	return "messages"
}
