import axios from 'axios'
import DashboardBundle from '../interfaces/Dashboard'

const url = `${process.env.REACT_APP_BACKEND_API}/totals`

export const getDashboard = (userId: string, year: number, month: string) => {
    return axios.post<DashboardBundle>(`${url}/dashboard`, {
        userId,
        year,
        month,
    })
}
