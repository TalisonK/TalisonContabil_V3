import axios from "axios"


export const getCard = async (id: string) => {
    const url = `${process.env.REACT_APP_BACKEND_API}/creditcard/${id}`
    return axios.get(url)
}