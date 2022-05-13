package entity

import (
	"time"
)

type Message struct {
	MessageId int       `json:"message_id,omitempty"`
	Message   string    `json:"message,omitempty"`
	TargetId  string    `json:"target_id,omitempty"`
	Time      time.Time `json:"time"`
	Action    string    `json:"action,omitempty"`
	StudentId int       `json:"student_id,omitempty"`
}
