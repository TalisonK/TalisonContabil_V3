package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "gorm.io/gorm"
)

type List struct {
	ID          string  `json:"id" gorm:"type:varchar(255);primary_key;"`
	ItemName    string  `json:"name" gorm:"type:varchar(255);not null;idx_list"`
	ItemValue   float64 `json:"value" gorm:"type:float;not null;idx_list"`
	ExpenseID   string  `json:"expenseID" gorm:"type:varchar(255);not null;idx_list"`
	ExpenseName string  `json:"expenseName" gorm:"type:varchar(255);not null;idx_list"`
	Expense     Expense `json:"expense" gorm:"constraint;"`
}

type ListDTO struct {
	ID          string  `json:"id"`
	ItemName    string  `json:"itemName"`
	ItemValue   float64 `json:"itemValue"`
	ExpenseID   string  `json:"expenseID"`
	ExpenseName string  `json:"expenseName"`
}

func (l *ListDTO) ToEntity() List {
	return List{
		ID:          l.ID,
		ItemName:    l.ItemName,
		ItemValue:   l.ItemValue,
		ExpenseID:   l.ExpenseID,
		ExpenseName: l.ExpenseName,
	}
}

func (l *List) ToDTO() ListDTO {
	return ListDTO{
		ID:          l.ID,
		ItemName:    l.ItemName,
		ItemValue:   l.ItemValue,
		ExpenseID:   l.ExpenseID,
		ExpenseName: l.ExpenseName,
	}
}

func (l *List) ToPrim() primitive.M {
	return primitive.M{
		"id":          l.ID,
		"itemName":    l.ItemName,
		"itemValue":   l.ItemValue,
		"expenseID":   l.ExpenseID,
		"expenseName": l.ExpenseName,
	}
}
