package domain

import (
	"time"

	"github.com/TalisonK/TalisonContabil/internal/constants"
	"github.com/TalisonK/TalisonContabil/pkg/timeHandler"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "gorm.io/gorm"
)

type Expense struct {
	ID            string     `json:"id" gorm:"type:varchar(255);primary_key;"`
	Description   string     `json:"description" gorm:"type:varchar(255);not null;idx_expense;"`
	PaymentMethod string     `json:"paymentMethod" gorm:"type:varchar(255);not null;idx_expense;"`
	Value         float64    `json:"value" gorm:"type:float;not null;idx_expense;"`
	CreatedAt     string     `json:"createdAt" gorm:"type:varchar(255);not null;idx_expense;"`
	UpdatedAt     string     `json:"updatedAt" gorm:"type:varchar(255);not null;idx_expense;"`
	PaidAt        string     `json:"paidAt" gorm:"type:varchar(255);not null;idx_expense;"`
	EndsAt        string     `json:"endsAt" gorm:"type:varchar(255);not null;idx_expense;"`
	UserID        string     `json:"userID" gorm:"type:varchar(255);not null;idx_expense;"`
	User          User       `json:"user" gorm:"constraint;"`
	CategoryID    string     `json:"categoryID" gorm:"type:varchar(255);not null;idx_expense;"`
	Category      Category   `json:"category" gorm:"constraint;"`
	CreditCardID  string     `json:"creditCardID" gorm:"type:varchar(255);not null;idx_expense;"`
	CreditCard    CreditCard `json:"creditCard" gorm:"constraint;"`
}

type ExpenseDTO struct {
	ID            string  `json:"id"`
	Description   string  `json:"description"`
	PaymentMethod string  `json:"paymentMethod"`
	Value         float64 `json:"value"`
	CreatedAt     string  `json:"createdAt"`
	UpdatedAt     string  `json:"updatedAt"`
	PaidAt        string  `json:"paidAt"`
	EndsAt        string  `json:"endsAt"`
	ActualParcel  int32   `json:"actualParcel"`
	TotalParcel   int32   `json:"totalParcel"`
	UserID        string  `json:"userID"`
	CategoryID    string  `json:"categoryID"`
	CreditCardID  string  `json:"creditCardID"`
	CategoryName  string  `json:"categoryName"`
	List          []List  `json:"list"`
}

func (e *ExpenseDTO) ToEntity() Expense {
	return Expense{
		ID:            e.ID,
		Description:   e.Description,
		PaymentMethod: e.PaymentMethod,
		Value:         e.Value,
		CreatedAt:     e.CreatedAt,
		UpdatedAt:     e.UpdatedAt,
		PaidAt:        e.PaidAt,
		UserID:        e.UserID,
		CategoryID:    e.CategoryID,
		CreditCardID:  e.CreditCardID,
		EndsAt:        e.EndsAt,
	}
}

func (e *Expense) ToDTO(searchDate string) ExpenseDTO {

	ac := int32(timeHandler.MonsthsAfterDate(e.PaidAt, searchDate)) + 1
	tp := int32(timeHandler.MonsthsAfterDate(e.PaidAt, e.EndsAt)) + 1

	return ExpenseDTO{
		ID:            e.ID,
		Description:   e.Description,
		PaymentMethod: e.PaymentMethod,
		Value:         e.Value,
		CreatedAt:     e.CreatedAt,
		UpdatedAt:     e.UpdatedAt,
		PaidAt:        e.PaidAt,
		UserID:        e.UserID,
		CategoryID:    e.CategoryID,
		CreditCardID:  e.CreditCardID,
		EndsAt:        e.EndsAt,
		ActualParcel:  ac,
		TotalParcel:   tp,
	}
}

func (e *Expense) ToPrim() primitive.M {
	pexp := primitive.M{}

	if e.ID != "" {
		id, _ := primitive.ObjectIDFromHex(e.ID)
		pexp["_id"] = id
	}

	pexp["description"] = e.Description
	pexp["paymentMethod"] = e.PaymentMethod
	pexp["value"] = e.Value

	createdAt, err := time.Parse(time.DateTime, e.CreatedAt)
	if err != nil {
		createdAt = time.Now()
	}
	pexp["createdAt"] = primitive.NewDateTimeFromTime(createdAt)
	updatedAt, err := time.Parse(time.DateTime, e.UpdatedAt)

	if err != nil {
		updatedAt = createdAt
	}
	pexp["updatedAt"] = updatedAt

	paidAt, err := time.Parse(time.DateTime, e.PaidAt)
	if err != nil {
		paidAt = time.Now()
	}
	pexp["paidAt"] = paidAt

	endsAt, err := time.Parse(time.DateTime, e.EndsAt)
	if err != nil {
		endsAt = paidAt
	}
	pexp["endsAt"] = endsAt

	pexp["userID"] = e.UserID
	pexp["categoryID"] = e.CategoryID
	pexp["creditCardID"] = e.CreditCardID

	return pexp
}

func (e *Expense) ToActivity() Activity {
	return Activity{
		ID:            e.ID,
		Description:   e.Description,
		PaymentMethod: e.PaymentMethod,
		Value:         e.Value,
		Type:          constants.EXPENSE,
		CreatedAt:     e.CreatedAt,
		UpdatedAt:     e.UpdatedAt,
		ActivityDate:  e.PaidAt,
		UserID:        e.UserID,
		CategoryID:    e.CategoryID,
		CreditCardID:  e.CreditCardID,
		EndsAt:        e.EndsAt,
		ActualParcel:  int32(timeHandler.MonthsAfterNow(e.PaidAt)),
		TotalParcel:   int32(timeHandler.MonsthsAfterDate(e.PaidAt, e.EndsAt)),
	}
}

func (e *ExpenseDTO) ToActivity() Activity {
	return Activity{
		ID:            e.ID,
		Description:   e.Description,
		PaymentMethod: e.PaymentMethod,
		Value:         e.Value,
		Type:          constants.EXPENSE,
		CreatedAt:     e.CreatedAt,
		UpdatedAt:     e.UpdatedAt,
		ActivityDate:  e.PaidAt,
		UserID:        e.UserID,
		CategoryID:    e.CategoryID,
		CreditCardID:  e.CreditCardID,
		CategoryName:  e.CategoryName,
		EndsAt:        e.EndsAt,
		ActualParcel:  e.ActualParcel,
		TotalParcel:   e.TotalParcel,
	}
}

func PrimToExpense(pexp primitive.M) Expense {
	exp := Expense{}

	if pexp["_id"] != nil {
		exp.ID = pexp["_id"].(primitive.ObjectID).Hex()
	}

	exp.Description = pexp["description"].(string)
	exp.PaymentMethod = pexp["paymentMethod"].(string)
	exp.Value = pexp["value"].(float64)
	exp.CreatedAt = pexp["createdAt"].(string)
	exp.UpdatedAt = pexp["updatedAt"].(string)
	exp.PaidAt = pexp["paidAt"].(string)
	exp.EndsAt = pexp["endsAt"].(string)
	exp.UserID = pexp["userID"].(string)
	exp.CategoryID = pexp["categoryID"].(string)
	exp.CreditCardID = pexp["creditCardID"].(string)

	return exp
}

func PrimToExpenseDto(pexp primitive.M) ExpenseDTO {
	exp := ExpenseDTO{}

	if pexp["_id"] != nil {
		exp.ID = pexp["_id"].(primitive.ObjectID).Hex()
	}

	exp.Description = pexp["description"].(string)
	exp.PaymentMethod = pexp["paymentMethod"].(string)
	exp.Value = pexp["value"].(float64)
	exp.CreatedAt = pexp["createdAt"].(string)
	exp.UpdatedAt = pexp["updatedAt"].(string)
	exp.PaidAt = pexp["paidAt"].(string)
	exp.EndsAt = pexp["endsAt"].(string)
	exp.UserID = pexp["userID"].(string)
	exp.CategoryID = pexp["categoryID"].(string)
	exp.CreditCardID = pexp["creditCardID"].(string)

	if pexp["categoryName"] != nil {
		exp.CategoryName = pexp["categoryName"].(string)
	}

	return exp
}
