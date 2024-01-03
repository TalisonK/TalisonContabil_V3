import Activity from "./Activity"
import Totals from "./Totals"

interface DashboardBundle {
    userId: string,
    year: string,
    month: string,
    updatedAt: Date,
    timeline: Activity[],
    resumes: {
        incomes: Totals[],
        expenses: Totals[],
        balances: Totals[]
    },
    charts: {
        incomeXexpense: {
            incomes: Totals[],
            expenses: Totals[]
        }
    }
}

export default DashboardBundle;