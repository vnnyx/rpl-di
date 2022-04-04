package entity

type CodeMessageEntity struct {
	Id      int    `gorm:"primaryKey;column:codemessage_id;autoIncrement"`
	Code    string `gorm:"column:code;"`
	Message string `gorm:"column:message"`
}

func (CodeMessageEntity) TableName() string {
	return "code_message"
}
