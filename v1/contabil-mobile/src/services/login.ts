import axios from "axios";
import api from "./api";
//import {BACKENDURL} from 'react-native-dotenv';


interface LoginResponse{
    token:string,
    expiresIn:number,
    userId:string
}

export const ping = () => {
    return api("/","get", null)
    .then((retorno) => {
        return retorno.status === 200;
    })
    .catch((e)=>{
        return false;
    });
}


export const logar = async(nome:string, senha:string):Promise<LoginResponse> => {

    const retorno = await api("/user/login","post", {
        "nome": nome,
        "senha": senha
    })

    if(retorno.status == 200){
        return retorno.data as LoginResponse;
    }
    else{
        return {"token":""} as LoginResponse;
    }
}