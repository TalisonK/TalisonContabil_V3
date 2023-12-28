import {Activity} from "./timeline.interface";
import {Totals} from "./Totals.interface";


export interface DashboardInterface {

    "userId": string,
    "year": string,
    "month": string,
    "updatedAt": Date,
    "timeline": Activity[],
    "resumes": {
        "incomes": Totals[],
        "expenses": Totals[],
        "balances": Totals[]
    },
    "charts": {
        "incomeXexpense": {
            "incomes": Totals[],
            "expenses": Totals[]
        }
    }
}
