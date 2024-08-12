package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreditCard struct {
	ID        string `json:"id" gorm:"type:varchar(255);primary_key;"`
	Flag      string `json:"flag" gorm:"type:varchar(255);not null;unique;idx_creditcard"`
	Bank      string `json:"bank" gorm:"type:varchar(255);not null;idx_creditcard"`
	ClosesAt  int    `json:"closesAt" gorm:"type:int;not null;idx_creditcard"`
	ExpiresAt int    `json:"expiresAt" gorm:"type:int;not null;idx_creditcard"`
	UserID    string `json:"userID" gorm:"type:varchar(255);not null;idx_creditcard"`
	User      User   `json:"user" gorm:"constraint;"`
	CreatedAt string `json:"createdAt" gorm:"type:varchar(255);not null;idx_creditcard"`
	UpdatedAt string `json:"updatedAt" gorm:"type:varchar(255);not null;idx_creditcard"`
}

func (c *CreditCard) ToPrim() primitive.M {
	pinc := primitive.M{}

	if c.ID != "" {
		id, _ := primitive.ObjectIDFromHex(c.ID)
		pinc["_id"] = id
	}

	pinc["flag"] = c.Flag
	pinc["bank"] = c.Bank
	pinc["closesAt"] = c.ClosesAt
	pinc["expiresAt"] = c.ExpiresAt
	pinc["userID"] = c.UserID
	pinc["createdAt"] = c.CreatedAt
	pinc["updatedAt"] = c.UpdatedAt

	return pinc
}

func PrimToCreditCard(prim primitive.M) CreditCard {
	return CreditCard{
		ID:        prim["_id"].(primitive.ObjectID).Hex(),
		Flag:      prim["flag"].(string),
		Bank:      prim["bank"].(string),
		ClosesAt:  prim["closesAt"].(int),
		ExpiresAt: prim["expiresAt"].(int),
		UserID:    prim["userID"].(string),
		CreatedAt: prim["createdAt"].(string),
		UpdatedAt: prim["updatedAt"].(string),
	}
}
