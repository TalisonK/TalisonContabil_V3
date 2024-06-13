import {api} from '../services/api';
import { Activity } from '../interfaces/Activity.interface';

interface ValueReturn{
    valor:number
}

export const getIncomeByMonth = async (month:Date):Promise<ValueReturn> => {
    const retorno = await api("/activity/get-income","post", {data: month.toISOString()});

    if(retorno.status === 200 && retorno.data !== null){
        return retorno.data as ValueReturn;
    }
    else{
        return {valor: 0.0};
    }
}

export const getExpenseByMonth = async (month:Date):Promise<ValueReturn> => {
    const retorno = await api("/activity/get-expense","post", {data: month.toISOString()});

    if(retorno.status === 200 && retorno.data !== null){
        return retorno.data as ValueReturn;
    }
    else{
        return {valor: 0.0};
    }
}

export const getActivitiesByMonth = async (month:Date) => {
    const retorno = await api("/activity/by-month","post", {data: month.toISOString()});

    let aux = retorno.data as Array<Activity>;

    aux = aux.sort((n1, n2) => {
        if(n1.dataPagamento > n2.dataPagamento){
            return -1;
        }
        if(n1.dataPagamento < n2.dataPagamento){
            return 1;
        }
        return 0;
    })
    
    return aux;
}

export const sendActivity = async (data:any) => {
    
    const retorno = await api("/activity/add", 'post', {...data});

    console.log(retorno);
    console.log(data);
    
    

    return (retorno);
    
}

export const filterActivity =async (data:any) => {
    
    const retorno = await api("/activity/filtro", "post", {...data});

    return retorno;
}

export const adiantaActivity = async (id:string) => {
    const retorno = await api("/activity/adianta", "post", {id})
    return retorno;
}

export const deletaActivity = async(id:string) => {
    const retorno = await api("/activity", "delete",{id})
}