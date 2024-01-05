import axios from 'axios'

const url = `${process.env.REACT_APP_BACKEND_API}`

export const categoryList = async () => {
    const response = await axios.get(url + '/category/all')
    return response.data
}

export const submitActivity = async (data: any) => {
    let result = {}

    let type = data.type

    delete data.type

    if (type === 'Income') {
        result = await axios.post(url + '/income', data)
    } else {
        result = await axios.post(url + '/expense', data)
    }
    return result
}
