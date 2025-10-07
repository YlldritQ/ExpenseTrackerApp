package models

import (
	"time"
)

type Expense struct {
	ID 			uint      	`gorm:"primaryKey" json:"id"`
	Amount      float64   	`gorm:"not null" json:"amount"`
	UserID		*uint		`json:"user_id"`
	User        *User		`gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
	Category 	string		`gorm:"not null" json:"category"`
	Description string    	`json:"description"`
	Date  		time.Time  	`json:"date"`
	CreatedAt   time.Time  	`json:"created_at"`
	UpdatedAt   time.Time  	`json:"updated_at"`
}