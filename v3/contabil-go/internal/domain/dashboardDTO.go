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
	ID            string  `json:"id"`
	Description   string  `json:"description"`
	PaymentMethod string  `json:"paymentMethod"`
	Value         float64 `json:"value"`
	Type          string  `json:"type"`
	CreatedAt     string  `json:"createdAt"`
	UpdatedAt     string  `json:"updatedAt"`
	ActivityDate  string  `json:"activityDate"`
	ActualParcel  int32   `json:"actualParcel"`
	TotalParcel   int32   `json:"totalParcel"`
	UserID        string  `json:"userID"`
	UserName      string  `json:"userName"`
	CategoryID    string  `json:"categoryID"`
	CategoryName  string  `json:"categoryName"`
	EndsAt        string  `json:"endsAt"`
}

type Resume struct {
	Actual  float64 `json:"actual"`
	Pass    float64 `json:"pass"`
	Balance float64 `json:"balance"`
}

func (a *Activity) ToIncomeDTO() IncomeDTO {
	return IncomeDTO{
		ID:          a.ID,
		Description: a.Description,
		Value:       a.Value,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
		ReceivedAt:  a.ActivityDate,
		UserID:      a.UserID,
	}
}

func (a *Activity) ToExpenseDTO() ExpenseDTO {
	return ExpenseDTO{
		ID:           a.ID,
		Description:  a.Description,
		Value:        a.Value,
		CreatedAt:    a.CreatedAt,
		UpdatedAt:    a.UpdatedAt,
		PaidAt:       a.ActivityDate,
		EndsAt:       a.EndsAt,
		UserID:       a.UserID,
		CategoryID:   a.CategoryID,
		CategoryName: a.CategoryName,
		ActualParcel: a.ActualParcel,
		TotalParcel:  a.TotalParcel,
	}
}
