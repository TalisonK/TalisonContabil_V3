package domain

type DashboardPacket struct {
	UserID            string                `json:"userID"`
	Month             string                `json:"month"`
	Year              int                   `json:"year"`
	UpdatedAt         string                `json:"updatedAt"`
	Timeline          []Activity            `json:"timeline"`
	Resumes           map[string]Resume     `json:"resumes"`
	ExpenseByCategory map[string]float64    `json:"expenseByCategory"`
	IncomevsExpense   []IncomevsExpense     `json:"incomeVSexpense"`
	ExpenseByMethod   map[string]float64    `json:"expenseByMethod"`
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
	ActualParcel  int32   `json:"actualParcel" gorm:"type:int;not null;idx_expense;"`
	TotalParcel   int32   `json:"totalParcel" gorm:"type:int;not null;idx_expense;"`
	UserID        string  `json:"userID" gorm:"type:varchar(255);not null;idx_expense;"`
	UserName      string  `json:"userName" gorm:"type:varchar(255);not null;idx_expense;"`
	CategoryID    string  `json:"categoryID" gorm:"type:varchar(255);not null;idx_expense;"`
	CategoryName  string  `json:"categoryName" gorm:"type:varchar(255);not null;idx_expense;"`
}

type Resume struct {
	Actual  float64 `json:"actual"`
	Pass    float64 `json:"pass"`
	Balance float64 `json:"balance"`
}
