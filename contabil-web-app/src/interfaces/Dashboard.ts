import Activity from './Activity'
import IncomeVSExpense from './IncomeVSExpense'
import Totals from './Totals'

interface DashboardBundle {
    userId: string
    year: string
    month: string
    updatedAt: Date
    timeline: Activity[]
    resumes: {
        incomes: Totals[]
        expenses: Totals[]
        balances: Totals[]
    }
    incomeVSexpense: IncomeVSExpense[]
    expenseByCategory: any
    expenseByMethod: any
    fixatedExpenses: {
        contas: Activity[]
        streaming: Activity[]
    }
}

export default DashboardBundle
