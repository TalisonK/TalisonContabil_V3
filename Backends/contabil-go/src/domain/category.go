package domain

import (
	_ "gorm.io/gorm"
)

type Category struct {
	ID          string `json:"id" gorm:"type:varchar(255);primary_key;"`
	Name        string `json:"name" gorm:"type:varchar(255);not null;unique;idx_category"`
	Description string `json:"description" gorm:"type:varchar(255);idx_category"`
	CreatedAt   string `json:"createdAt" gorm:"type:varchar(255);not null;idx_user"`
	UpdatedAt   string `json:"updatedAt" gorm:"type:varchar(255);not null;idx_user"`
}
