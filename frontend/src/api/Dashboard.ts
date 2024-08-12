import axios from 'axios'
import DashboardBundle from '../interfaces/Dashboard'

const url = `${process.env.REACT_APP_BACKEND_API}/total`

export const getDashboard = (userId: string, year: number, month: string) => {
    return axios.post<DashboardBundle>(`${url}/dashboard`, {
        userId,
        year,
        month,
    })
}

export const clearCache = () => {
    return axios.delete(`${url}/clear-cache`)
}