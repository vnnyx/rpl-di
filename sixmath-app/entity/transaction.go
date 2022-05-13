package entity

import (
	"time"
)

type Transaction struct {
	TransactionId   int       `json:"transaction_id,omitempty"`
	StudentId       int       `json:"student_id,omitempty"`
	TotalAmount     float64   `json:"total_amount,omitempty"`
	Subscription    bool      `json:"subscription,omitempty"`
	TransactionDate time.Time `json:"transaction_date"`
	Status          bool      `json:"status,omitempty"`
}
