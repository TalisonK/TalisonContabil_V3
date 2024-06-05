package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

type UserDTO struct {
	ID        string `json:"id" gorm:"type:varchar(255);primary_key;"`
	Name      string `json:"name" gorm:"type:varchar(255);not null;unique;idx_user"`
	Role      string `json:"role" gorm:"type:varchar(255);not null;idx_user"`
	CreatedAt string `json:"createdAt" gorm:"type:varchar(255);not null;idx_user"`
	UpdatedAt string `json:"updatedAt" gorm:"type:varchar(255);not null;idx_user"`
}

func (u *UserDTO) ToEntity() User {
	return User{
		ID:        u.ID,
		Name:      u.Name,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u *User) ToDTO() UserDTO {
	return UserDTO{
		ID:        u.ID,
		Name:      u.Name,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u *User) ToPrim() primitive.M {
	pinc := primitive.M{}

	if u.ID != "" {
		id, _ := primitive.ObjectIDFromHex(u.ID)
		pinc["_id"] = id
	}

	pinc["name"] = u.Name
	pinc["role"] = u.Role
	pinc["password"] = u.Password
	pinc["salt"] = u.Salt
	pinc["createdAt"] = u.CreatedAt
	pinc["updatedAt"] = u.UpdatedAt

	return pinc
}

func PrimToUser(user primitive.M) User {
	var usuario User

	usuario.ID = user["_id"].(primitive.ObjectID).Hex()
	usuario.Name = user["name"].(string)
	usuario.Password = user["password"].(string)
	usuario.Role = user["role"].(string)
	usuario.Salt = user["salt"].(string)
	usuario.CreatedAt = user["createdAt"].(primitive.DateTime).Time().Format(time.RFC3339)
	usuario.UpdatedAt = user["updatedAt"].(primitive.DateTime).Time().Format(time.RFC3339)

	return usuario
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
