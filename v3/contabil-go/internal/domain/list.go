package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "gorm.io/gorm"
)

type List struct {
	ID        string  `json:"id" gorm:"type:varchar(255);primary_key;"`
	ItemName  string  `json:"name" gorm:"type:varchar(255);not null;idx_list"`
	ItemPrice float64 `json:"price" gorm:"type:float;not null;idx_list"`
	ExpenseID string  `json:"expenseID" gorm:"type:varchar(255);not null;idx_list"`
	Expense   Expense `json:"expense" gorm:"constraint;"`
}

type ListDTO struct {
	ID        string  `json:"id"`
	ItemName  string  `json:"itemName"`
	ItemPrice float64 `json:"itemValue"`
	ExpenseID string  `json:"expenseID"`
}

func (l *ListDTO) ToEntity() List {
	return List{
		ID:        l.ID,
		ItemName:  l.ItemName,
		ItemPrice: l.ItemPrice,
		ExpenseID: l.ExpenseID,
	}
}

func (l *List) ToDTO() ListDTO {
	return ListDTO{
		ID:        l.ID,
		ItemName:  l.ItemName,
		ItemPrice: l.ItemPrice,
		ExpenseID: l.ExpenseID,
	}
}

func (l *List) ToPrim() primitive.M {
	return primitive.M{
		"id":        l.ID,
		"itemName":  l.ItemName,
		"ItemPrice": l.ItemPrice,
		"expenseID": l.ExpenseID,
	}
}
