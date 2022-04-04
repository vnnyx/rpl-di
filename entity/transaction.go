package entity

import "time"

type Transaction struct {
	Id              int           `gorm:"column:transaction_id;primaryKey;autoIncrement"`
	Child           StudentEntity `gorm:"association_foreignKey:ChildId"`
	ChildId         int           `gorm:"column:child_id"`
	TotalAmount     float64       `gorm:"column:total_amount"`
	Subscription    int           `gorm:"column:subcription"`
	TransactionDate time.Time     `gorm:"column:transaction_date"`
	Status          bool          `gorm:"column:status"`
}

func (Transaction) TableName() string {
	return "transactions"
}
