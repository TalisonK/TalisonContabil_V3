import api from "./api"

interface Category{
    id:string,
    nome:string
}

export const getFilteredCategory = async (filtro:string):Promise<Array<Category>> => {

    const aux = await api("/categoria/filter", "post", {nome: filtro});

    return aux.data as Array<Category>;
}