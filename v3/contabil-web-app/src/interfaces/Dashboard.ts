import Activity from './Activity'
import IncomeVSExpense from './IncomeVSExpense'
import Totals from './Totals'
import ResumeBundle from './Resume'

interface DashboardBundle {
    userId: string
    year: string
    month: string
    updatedAt: Date
    timeline: Activity[]
    resumes: {
        income: ResumeBundle
        expense: ResumeBundle
        balance: ResumeBundle
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
