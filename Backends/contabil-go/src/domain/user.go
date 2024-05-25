package domain

import (
	_ "gorm.io/gorm"
)

type User struct {
	ID        string `json:"id" gorm:"type:varchar(255);primary_key;"`
	Name      string `json:"name" gorm:"type:varchar(255);not null;unique;idx_user"`
	Password  string `json:"password" gorm:"type:varchar(255);not null;idx_user"`
	Salt      string `json:"salt" gorm:"type:varchar(255);not null;idx_user"`
	Role      string `json:"role" gorm:"type:varchar(255);not null;idx_user"`
	CreatedAt string `json:"createdAt" gorm:"type:varchar(255);not null;idx_user"`
	UpdatedAt string `json:"updatedAt" gorm:"type:varchar(255);not null;idx_user"`
}

/*
type User struct {
	RawData           map[string]interface{}
	ID        string `json:"id" gorm:"type:varchar(255);primary_key;"`
	Password  string `json:"password" gorm:"type:varchar(255);not null;idx_user"`
	Role      string `json:"role" gorm:"type:varchar(255);not null;idx_user"`
	CreatedAt string `json:"createdAt" gorm:"type:varchar(255);not null;idx_user"`
	UpdatedAt string `json:"updatedAt" gorm:"type:varchar(255);not null;idx_user"`
	Provider          string
	Email             string
	Name              string
	FirstName         string
	LastName          string
	NickName          string
	Description       string
	UserID            string
	AvatarURL         string
	Location          string
	AccessToken       string
	AccessTokenSecret string
	RefreshToken      string
	ExpiresAt         time.Time
	IDToken           string
}*/
