package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "gorm.io/gorm"
)

type Category struct {
	ID          string `json:"id" gorm:"type:varchar(255);primary_key;"`
	Name        string `json:"name" gorm:"type:varchar(255);not null;unique;idx_category"`
	Description string `json:"description" gorm:"type:varchar(255);idx_category"`
	CreatedAt   string `json:"createdAt" gorm:"type:varchar(255);not null;idx_user"`
	UpdatedAt   string `json:"updatedAt" gorm:"type:varchar(255);not null;idx_user"`
}

func (c *Category) ToPrim() primitive.M {
	pcat := primitive.M{}

	if c.ID != "" {
		id, _ := primitive.ObjectIDFromHex(c.ID)
		pcat["_id"] = id
	}

	pcat["name"] = c.Name
	pcat["description"] = c.Description

	createdAt, _ := time.Parse(time.DateTime, c.CreatedAt)
	pcat["createdAt"] = primitive.NewDateTimeFromTime(createdAt)

	updatedAt, _ := time.Parse(time.DateTime, c.UpdatedAt)
	pcat["updatedAt"] = primitive.NewDateTimeFromTime(updatedAt)

	return pcat
}

func PrimToCategory(prim primitive.M) Category {
	return Category{
		ID:          prim["_id"].(primitive.ObjectID).Hex(),
		Name:        prim["name"].(string),
		Description: prim["description"].(string),
		CreatedAt:   prim["createdAt"].(primitive.DateTime).Time().Format(time.DateTime),
		UpdatedAt:   prim["updatedAt"].(primitive.DateTime).Time().Format(time.DateTime),
	}
}
