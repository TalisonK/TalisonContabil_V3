package domain

import (
	_ "gorm.io/gorm"
)

type UserDTO struct {
	ID        string `json:"id" gorm:"type:varchar(255);primary_key;"`
	Name      string `json:"name" gorm:"type:varchar(255);not null;unique;idx_user"`
	Role      string `json:"role" gorm:"type:varchar(255);not null;idx_user"`
	CreatedAt string `json:"createdAt" gorm:"type:varchar(255);not null;idx_user"`
	UpdatedAt string `json:"updatedAt" gorm:"type:varchar(255);not null;idx_user"`
}
