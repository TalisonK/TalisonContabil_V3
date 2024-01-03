import axios from 'axios'
import DashboardBundle from '../interfaces/Dashboard'

const url = 'http://localhost:8080/api/totals'

export const getDashboard = (userId: string, year: string, month: string) => {
    return axios.post<DashboardBundle>(`${url}/dashboard`, {
        userId,
        year,
        month,
    })
}
