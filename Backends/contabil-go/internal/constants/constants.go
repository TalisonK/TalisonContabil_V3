package constants

const INCOME string = "Income"

const EXPENSE string = "Expense"

const CLOUD string = "Cloud"

const LOCAL string = "Local"

const ALL string = "All"

var payment_methods []string = []string{"CREDIT_CARD", "DEBIT_CARD", "PIX", "MONEY"}

func GetMethods() []string {
	return payment_methods
}
