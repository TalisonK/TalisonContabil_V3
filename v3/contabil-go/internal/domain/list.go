package domain

import (
	_ "gorm.io/gorm"
)

type List struct {
	ID        string  `json:"id" gorm:"type:varchar(255);primary_key;"`
	CreateAt  string  `json:"createAt" gorm:"type:varchar(255);not null;idx_list"`
	UpdateAt  string  `json:"updateAt" gorm:"type:varchar(255);not null;idx_list"`
	ItemName  string  `json:"itemName" gorm:"type:varchar(255);not null;idx_list"`
	ItemValue float64 `json:"itemValue" gorm:"type:float;not null;idx_list"`
	ExpenseID string  `json:"expenseID" gorm:"type:varchar(255);not null;idx_list"`
	Expense   Expense `json:"expense" gorm:"constraint;"`
}

type ListDTO struct {
	ID        string  `json:"id" gorm:"type:varchar(255);primary_key;"`
	CreateAt  string  `json:"createAt" gorm:"type:varchar(255);not null;idx_list"`
	UpdateAt  string  `json:"updateAt" gorm:"type:varchar(255);not null;idx_list"`
	ItemName  string  `json:"itemName" gorm:"type:varchar(255);not null;idx_list"`
	ItemValue float64 `json:"itemValue" gorm:"type:float;not null;idx_list"`
	ExpenseID string  `json:"expenseID" gorm:"type:varchar(255);not null;idx_list"`
}

func (l *ListDTO) ToEntity() List {
	return List{
		ID:        l.ID,
		CreateAt:  l.CreateAt,
		UpdateAt:  l.UpdateAt,
		ItemName:  l.ItemName,
		ItemValue: l.ItemValue,
		ExpenseID: l.ExpenseID,
	}
}

func (l *List) ToDTO() ListDTO {
	return ListDTO{
		ID:        l.ID,
		CreateAt:  l.CreateAt,
		UpdateAt:  l.UpdateAt,
		ItemName:  l.ItemName,
		ItemValue: l.ItemValue,
		ExpenseID: l.ExpenseID,
	}
}
