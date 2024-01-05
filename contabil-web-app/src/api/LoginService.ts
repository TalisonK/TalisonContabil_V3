import axios from 'axios'

const url = `${process.env.REACT_APP_BACKEND_API}/user`

export const login = (name: string, password: string): any => {
    return axios.post(url + '/login', {
        name: name,
        password: password,
    })
}
