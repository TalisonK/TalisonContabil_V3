package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "gorm.io/gorm"
)

type IncomeSlice []Income

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

type IncomeDTO struct {
	ID          string  `json:"id" gorm:"type:varchar(255);primary_key;"`
	Description string  `json:"description" gorm:"type:varchar(255);not null;idx_income"`
	Value       float64 `json:"value" gorm:"type:float;not null;idx_income"`
	CreatedAt   string  `json:"createdAt" gorm:"type:varchar(255);not null;idx_income"`
	UpdatedAt   string  `json:"updatedAt" gorm:"type:varchar(255);not null;idx_income"`
	ReceivedAt  string  `json:"receivedAt" gorm:"type:varchar(255);not null;idx_income"`
	UserID      string  `json:"userID" gorm:"type:varchar(255);not null;idx_income"`
}

func (i *IncomeDTO) ToEntity() Income {
	return Income{
		ID:          i.ID,
		Description: i.Description,
		Value:       i.Value,
		CreatedAt:   i.CreatedAt,
		UpdatedAt:   i.UpdatedAt,
		ReceivedAt:  i.ReceivedAt,
		UserID:      i.UserID,
	}
}

func (i *Income) ToDTO() IncomeDTO {
	return IncomeDTO{
		ID:          i.ID,
		Description: i.Description,
		Value:       i.Value,
		CreatedAt:   i.CreatedAt,
		UpdatedAt:   i.UpdatedAt,
		ReceivedAt:  i.ReceivedAt,
		UserID:      i.UserID,
	}
}

func (i *IncomeDTO) ToPrim() primitive.M {
	pinc := primitive.M{}

	if i.ID != "" {
		id, _ := primitive.ObjectIDFromHex(i.ID)
		pinc["_id"] = id
	}

	pinc["description"] = i.Description
	pinc["value"] = i.Value
	createdAt, err := time.Parse(time.RFC3339, i.CreatedAt)
	if err != nil {
		createdAt = time.Now()
	}
	pinc["createdAt"] = primitive.NewDateTimeFromTime(createdAt)
	updatedAt, err := time.Parse(time.RFC3339, i.UpdatedAt)

	if err != nil {
		updatedAt = createdAt
	}
	pinc["updatedAt"] = updatedAt

	receivedAt, err := time.Parse(time.RFC3339, i.ReceivedAt)
	if err != nil {
		receivedAt = time.Now()
	}
	pinc["receivedAt"] = receivedAt
	pinc["userId"] = i.UserID

	return pinc
}

func PrimToIncome(income primitive.M) *Income {
	newIncome := Income{}

	newIncome.ID = income["_id"].(primitive.ObjectID).Hex()

	if value, ok := income["value"].(int32); ok {
		newIncome.Value = float64(value)
	} else {
		newIncome.Value = income["value"].(float64)
	}

	newIncome.Description = income["description"].(string)
	newIncome.ReceivedAt = income["receivedAt"].(primitive.DateTime).Time().Format(time.RFC3339)
	newIncome.UserID = income["userId"].(string)
	newIncome.CreatedAt = income["createdAt"].(primitive.DateTime).Time().Format(time.RFC3339)
	if income["updatedAt"] == nil {
		newIncome.UpdatedAt = income["createdAt"].(primitive.DateTime).Time().Format(time.RFC3339)
	} else {
		newIncome.UpdatedAt = income["updatedAt"].(primitive.DateTime).Time().Format(time.RFC3339)
	}

	return &newIncome
}

func (i IncomeSlice) DTOs() []IncomeDTO {
	incomes := []IncomeDTO{}

	for _, income := range i {
		incomes = append(incomes, income.ToDTO())
	}

	return incomes
}
