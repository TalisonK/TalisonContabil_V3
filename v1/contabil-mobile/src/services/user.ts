import {api} from '../services/api';

export const getAllUser = async(signOut: any) => {

    const retorno = await api("/user", "get", null, signOut);

    return retorno;

}

export const createUser = async(nome: string, senha: string, signOut: any) => {

    const retorno = await api("/user", 'post', {nome, senha}, signOut);
    return retorno.data;

}

export const editUser = async(id: string, nome:string, senha:string, signOut: any) => {

    const retorno = await api("/user", 'put', {id,nome, senha}, signOut);
    return retorno.data;

}