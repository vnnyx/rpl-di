package entity

import (
	"time"
)

type TransactionEntity struct {
	TransactionId   int         `gorm:"column:transaction_id;primaryKey;autoIncrement;type:int"`
	Student         *UserEntity `gorm:"association_foreignkey:UserId"`
	StudentId       int         `gorm:"column:student_id;type:int"`
	TotalAmount     float64     `gorm:"column:total_amount;type:double"`
	Subscription    bool        `gorm:"column:subscription;type:bool"`
	TransactionDate time.Time   `gorm:"column:transaction_date"`
	Status          bool        `gorm:"column:status;type:bool"`
}

func (TransactionEntity) TableName() string {
	return "transaction"
}
