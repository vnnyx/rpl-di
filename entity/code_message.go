package entity

type CodeMessageEntity struct {
	Id      int    `gorm:"column:codemessage_id;primaryKey;int;autoIncrement"`
	Code    string `gorm:"column:code;type:text"`
	Message string `gorm:"column:message"`
}

func (CodeMessageEntity) TableName() string {
	return "code_message"
}
