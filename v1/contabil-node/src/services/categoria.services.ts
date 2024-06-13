
const { prisma } = require("../../prisma/index")
const { Categorias } = require("@prisma/client")


interface ListReturn{
    id:string,
    nome:string
}
export default class CategoriaService {

    

    static getCategorias = async () => {

        const aux = await prisma.categoria.findMany({
            orderBy: {
                nome: 'asc'
            }
        })
        return aux
    }

    static getFilteredCategorias = async (filter: string) => {
        return this.getCategorias()
            .then((aux:Array<ListReturn>) => {
                const list = []
                const a: string = "";

                for (let i of aux) {
                    if (i.nome.toLocaleLowerCase().includes(filter.toLocaleLowerCase())) {
                        list.push(i)
                    }
                }
                return list;
            })
            .catch(() =>([]))
    }
}








