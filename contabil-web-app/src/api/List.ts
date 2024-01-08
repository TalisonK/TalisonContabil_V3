import axios from 'axios'
import Activity from '../interfaces/Activity'

const url = `${process.env.REACT_APP_BACKEND_API}`

export const getList = async (id: string) => {
    const responseExpense = await axios.get<Activity[]>(
        url + '/expense/all/' + id
    )
    const responseIncome = await axios.get<Activity[]>(
        url + '/income/all/' + id
    )

    return [...responseExpense.data, ...responseIncome.data]
}
