package entity

import "time"

type Transaction struct {
	Id              int
	ChildId         StudentEntity
	TotalAmount     float64
	Subscription    int
	TransactionDate time.Time
	Status          bool
}
