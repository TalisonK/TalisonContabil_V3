import axios from "axios"



export const login = (name: string, password: string): any => {

    return axios.post("http://localhost:8080/api/user/login", {
        name: name,
        password: password
    })
}
