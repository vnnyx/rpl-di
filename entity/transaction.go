package entity

import "time"

type Transaction struct {
	TransactionId   int            `gorm:"column:transaction_id;primaryKey;autoIncrement;type:int"`
	Child           *StudentEntity `gorm:"association_foreignKey:StudentId"`
	ChildId         int            `gorm:"column:child_id;type:int"`
	TotalAmount     float64        `gorm:"column:total_amount;type:double"`
	Subscription    int            `gorm:"column:subscription;type:bool"`
	TransactionDate time.Time      `gorm:"column:transaction_date"`
	Status          bool           `gorm:"column:status;type:bool"`
}

func (Transaction) TableName() string {
	return "transaction"
}
