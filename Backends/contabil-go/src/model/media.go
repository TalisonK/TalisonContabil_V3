package model

import (
	_ "gorm.io/gorm"
)


type Media struct {
	ID          string    `gorm:"type:char(36);primary_key;"`
	Name        string    `json:"name" gorm:"type:varchar(255);not null;unique;idx_media"`
	MimeType    string    `json:"name" gorm:"type:varchar(255);not null;idx_media"`
	Size        int64     `gorm:"type:bigint;not null;not null;idx_media"`
	StoragePath string    `gorm:"type:varchar(255);not null"`
}