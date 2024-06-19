import axios from 'axios'
import Activity from '../interfaces/Activity'

const url = `${process.env.REACT_APP_BACKEND_API}`

export const getList = async (id: string) => {
    
    const response = await axios.get<Activity[]>(
        url + '/total/activities/' + id
    )

    return response.data
}

export const deleteActivity = async (bucket: Activity[]) => {
    
    const response = await axios.post<boolean>(url + '/total/bucket', bucket)

    return response
}
