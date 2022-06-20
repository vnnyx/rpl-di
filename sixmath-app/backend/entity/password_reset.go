package entity

import (
	"time"
)

type PasswordReset struct {
	UserEmail string    `gorm:"column_name:email;type:varchar(100)" json:"email"`
	User      *User     `gorm:"references:Email" json:"-"`
	Token     int       `gorm:"column_name:token;type:bigint" json:"token"`
	ExpiredAt time.Time `gorm:"column_name:expired_at;type:datetime" json:"expired_at"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"-"`
}

func (PasswordReset) TableName() string {
	return "password_resets"
}
