package model

import (
	_ "gorm.io/gorm"
)

type Income struct {
	ID          string  `json:"id" gorm:"type:varchar(255);primary_key;"`
	Description string  `json:"description" gorm:"type:varchar(255);not null;idx_income"`
	Value       float64 `json:"value" gorm:"type:float;not null;idx_income"`
	CreatedAt   string  `json:"createdAt" gorm:"type:varchar(255);not null;idx_income"`
	UpdatedAt   string  `json:"updatedAt" gorm:"type:varchar(255);not null;idx_income"`
	ReceivedAt  string  `json:"receivedAt" gorm:"type:varchar(255);not null;idx_income"`
	UserID      string  `json:"userID" gorm:"type:varchar(255);not null;idx_income"`
	User        User    `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
