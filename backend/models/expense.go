package models

import (
	"time"
)

type Expense struct {
	ID 			uint      	`gorm:"primaryKey" json:"id"`
	Amount      float64   	`json:"amount"`
	Category 	string		`json:"category"`
	Description string    	`json:"description"`
	Date  		time.Time  	`json:"date"`
	CreatedAt   time.Time  	`json:"created_at"`
	UpdatedAt   time.Time  	`json:"updated_at"`
}