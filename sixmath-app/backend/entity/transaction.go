package entity

import (
	"time"
)

type Transaction struct {
	TransactionId   int       `gorm:"column:transaction_id;primaryKey;autoIncrement;type:int" json:"transaction_id,omitempty"`
	Student         *User     `gorm:"association_foreignkey:UserId" json:"-"`
	StudentId       int       `gorm:"column:student_id;type:int" json:"student_id,omitempty"`
	TotalAmount     float64   `gorm:"column:total_amount;type:double" json:"total_amount,omitempty"`
	Subscription    bool      `gorm:"column:subscription;type:bool" json:"subscription,omitempty"`
	TransactionDate time.Time `gorm:"column:transaction_date" json:"transaction_date"`
	Status          bool      `gorm:"column:status;type:bool" json:"status,omitempty"`
}

func (Transaction) TableName() string {
	return "transactions"
}
