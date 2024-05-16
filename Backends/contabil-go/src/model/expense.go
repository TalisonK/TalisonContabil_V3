package model

import (
	_ "gorm.io/gorm"
)

type Expense struct {
	ID            string   `json:"id" gorm:"type:varchar(255);primary_key;"`
	Description   string   `json:"description" gorm:"type:varchar(255);not null;idx_expense;"`
	PaymentMethod string   `json:"paymentMethod" gorm:"type:varchar(255);not null;idx_expense;"`
	Value         float64  `json:"value" gorm:"type:float;not null;idx_expense;"`
	CreatedAt     string   `json:"createdAt" gorm:"type:varchar(255);not null;idx_expense;"`
	UpdatedAt     string   `json:"updatedAt" gorm:"type:varchar(255);not null;idx_expense;"`
	PaidAt        string   `json:"paidAt" gorm:"type:varchar(255);not null;idx_expense;"`
	UserID        string   `json:"userID" gorm:"type:varchar(255);not null;idx_expense;"`
	User          User     `json:"user" gorm:"constraint;"`
	CategoryID    string   `json:"categoryID" gorm:"type:varchar(255);not null;idx_expense;"`
	Category      Category `json:"category" gorm:"constraint;"`
}
