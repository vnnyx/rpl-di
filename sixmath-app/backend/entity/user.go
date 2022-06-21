package entity

import "time"

type User struct {
	UserId          int       `gorm:"column:user_id;primaryKey;autoIncrement" json:"user_id,omitempty"`
	Username        string    `gorm:"column:username;unique;type:varchar(50)" json:"username,omitempty"`
	Email           string    `gorm:"column:email;type:varchar(100);unique" json:"email,omitempty"`
	Handphone       string    `gorm:"column:handphone;type:varchar(20)" json:"handphone,omitempty"`
	Password        string    `gorm:"column:password;type:varchar(255)" json:"password,omitempty"`
	Role            string    `gorm:"column:role;type:varchar(10)" json:"role,omitempty"`
	Avatar          string    `gorm:"column:avatar;type:mediumtext" json:"avatar,omitempty"`
	Certificate     string    `gorm:"column:certificate;type:mediumtext" json:"certificate,omitempty"`
	StudentUsername string    `gorm:"column:student_username;type:varchar(50)" json:"student_username,omitempty"`
	Age             int       `gorm:"column:age;type:int" json:"age,omitempty"`
	Description     string    `gorm:"column:description;type:varchar(255)" json:"description,omitempty"`
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime" json:"-"`
}

func (User) TableName() string {
	return "users"
}
