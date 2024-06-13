import axios from "axios";
import AsyncStorage from "@react-native-community/async-storage";
import { useAuth } from "../contexts/auth";

interface Retorno{
    message?: string;
    status: number,
    data: object | null
}

const client = axios.create({
    //baseURL: 'https://talison-contabil.onrender.com'
    baseURL: 'http://10.0.0.108:3333'
})

export const api = async (url:string, method:string, data:object | null, signOut?:any): Promise<Retorno> => {

    const token = JSON.parse(await AsyncStorage.getItem('@Auth:token') as string);
    
    if(data){
        return await client({
            url,
            method,
            data,
            headers: {'authorization': `Bearer ${token}`}
        })
        .then(retorno => ({data: retorno.data, status: retorno.status}))
        .catch(e => {
            //signOut(); 
            return e as Retorno}
        )
    }

    else{
        return await client({
            url,
            method,
            headers: {'authorization': `Bearer ${token}`}
        })
        .then(retorno => ({data: retorno.data, status: retorno.status}))
        .catch(e => {signOut(); return {} as Retorno})
    }
}


export default api;