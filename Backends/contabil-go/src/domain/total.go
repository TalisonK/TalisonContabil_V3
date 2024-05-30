package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "gorm.io/gorm"
)

type Total struct {
	ID         string  `json:"id" gorm:"type:varchar(255);primary_key;"`
	UserID     string  `json:"userID" gorm:"type:varchar(255);not null;idx_total"`
	CreatedAt  string  `json:"createdAt" gorm:"type:varchar(255);not null;idx_total"`
	UpdatedAt  string  `json:"updatedAt" gorm:"type:varchar(255);not null;idx_total"`
	TotalValue float64 `json:"totalValue" gorm:"type:float;not null;idx_total"`
	Month      string  `json:"month" gorm:"type:varchar(255);not null;idx_total"`
	Year       int32   `json:"year" gorm:"type:int;not null;idx_total"`
	Type       string  `json:"type" gorm:"type:varchar(255);not null;idx_total"`
}

func (t Total) ToPrim() primitive.M {
	pt := primitive.M{}

	if t.ID != "" {
		id, _ := primitive.ObjectIDFromHex(t.ID)
		pt["_id"] = id
	}

	pt["userID"] = t.UserID

	if t.CreatedAt == "" {
		createdAt, _ := time.Parse(time.RFC3339, t.CreatedAt)
		pt["createdAt"] = createdAt
	}

	if t.UpdatedAt == "" {
		updatedAt, _ := time.Parse(time.RFC3339, t.UpdatedAt)
		pt["updatedAt"] = updatedAt
	}
	pt["totalValue"] = t.TotalValue
	pt["month"] = t.Month
	pt["year"] = t.Year
	pt["type"] = t.Type

	return pt
}

func PrimToTotal(pt primitive.M) Total {
	t := Total{}

	if pt["_id"] != nil {
		t.ID = pt["_id"].(primitive.ObjectID).Hex()
	}

	t.UserID = pt["userID"].(string)

	if pt["createdAt"] != nil {
		t.CreatedAt = pt["createdAt"].(time.Time).Format(time.RFC3339)
	}

	if pt["updatedAt"] != nil {
		t.UpdatedAt = pt["updatedAt"].(time.Time).Format(time.RFC3339)
	}

	t.TotalValue = pt["totalValue"].(float64)
	t.Month = pt["month"].(string)
	year, _ := pt["year"].(int32)
	t.Year = year
	t.Type = pt["type"].(string)

	return t
}
