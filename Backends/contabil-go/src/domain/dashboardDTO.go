package domain

type DashboardPacket struct {
	UserID            string                `json:"userID"`
	Month             string                `json:"month"`
	Year              int                   `json:"year"`
	UpdatedAt         string                `json:"updatedAt"`
	Timeline          []Activity            `json:"timeline"`
	Resume            map[string][]Total    `json:"resume"`
	ExpensevsCategory map[string]float64    `json:"expensevsCategory"`
	ExpensevsMethod   map[string]float64    `json:"expensevsMethod"`
	FixatedExpenses   map[string][]Activity `json:"fixatedExpenses"`
}

type IncomevsExpense struct {
	Income  float64 `json:"income"`
	Expense float64 `json:"expense"`
	Month   string  `json:"month"`
	Year    int     `json:"year"`
}

type Activity struct {
	ID            string  `json:"id" gorm:"type:varchar(255);primary_key;"`
	Description   string  `json:"description" gorm:"type:varchar(255);not null;idx_expense;"`
	PaymentMethod string  `json:"paymentMethod" gorm:"type:varchar(255);not null;idx_expense;"`
	Value         float64 `json:"value" gorm:"type:float;not null;idx_expense;"`
	Type          string  `json:"type" gorm:"type:varchar(255);not null;idx_expense;"`
	CreatedAt     string  `json:"createdAt" gorm:"type:varchar(255);not null;idx_expense;"`
	UpdatedAt     string  `json:"updatedAt" gorm:"type:varchar(255);not null;idx_expense;"`
	ActivityDate  string  `json:"paidAt" gorm:"type:varchar(255);not null;idx_expense;"`
	UserID        string  `json:"userID" gorm:"type:varchar(255);not null;idx_expense;"`
	CategoryID    string  `json:"categoryID" gorm:"type:varchar(255);not null;idx_expense;"`
}
