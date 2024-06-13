import { Activity } from "@prisma/client"
import CategoriaService from "./categoria.services"

const { prisma } = require("../../prisma/index")

interface retorno {
    status: boolean,
    message?: string,
    body?: object
}

export default class ActivityServices {
    //buscar todas as compras
    static listAllActivities = async () => {
        const retorno = await prisma.activity.findMany({
            orderBy: [
                { createdAt: 'desc' }
            ]
        })
        return retorno
    }

    static listAllByMonth = async (entry: string) => {

        const atual: Date = new Date(entry);

        let filtroInicio: Date = new Date(atual.getFullYear(), atual.getMonth(), 1, -3);
        let filtroFinal: Date = new Date(atual.getFullYear(), atual.getMonth() + 1, 0, 20, 59, 59);

        const retorno = await prisma.activity.findMany({
            where: {
                dataPagamento: {
                    gte: filtroInicio,
                    lte: filtroFinal,
                },
            },
            orderBy: [
                {
                    createdAt: 'desc'
                }
            ],
        })
        return retorno;
    }

    static getFromMonth = async (entry: string, tipo: string): Promise<number> => {

        const lista = await this.listAllByMonth(entry);

        let sum = 0;

        for (let obj of lista) {
            if (obj.tipo === tipo) {
                sum += obj.valor;
            }
        }
        return sum;
    }


    //buscar compras por descrição

    static listAllByFilter = async (filtro: any) => {

        const aux: any = {}

        let cond = false;

        if (filtro["descricao"]) { aux['descricao'] = {contains: filtro["descricao"]}; cond = true }
        if (filtro["metodo"]) { aux['metodo'] = filtro["metodo"]; cond = true }
        if (filtro["categoria"]) { aux['categoria'] = filtro["categoria"]; cond = true }
        if (filtro["user"]) { aux['user'] = filtro["user"]; cond = true }
        if (filtro["valor"]) { aux['valor'] = filtro["valor"]; cond = true }
        if (filtro["tipo"]) { aux["tipo"] = filtro["tipo"]; cond = true }

        let data: any = {};

        if (filtro["dataInicio"] !== "") {
            data = { ...data, gte: new Date(filtro["dataInicio"]) }
            aux["dataPagamento"] = data;
            cond = true
        }
        if (filtro["dataFim"] !== "") {
            data = { ...data, 'lte': new Date(filtro["dataFim"]) }
            aux["dataPagamento"] = data;
            cond = true
        }
        let retorno = null       

        if (cond) {
            retorno = await prisma.activity.findMany({
                orderBy: [
                    { dataPagamento: 'desc' }
                ],
                where: aux
            })
            return retorno
        }
        else{
            retorno = await prisma.activity.findMany({
                orderBy: [
                    { dataPagamento: 'desc' }
                ]
            })
            return retorno
        }
    }

    //buscar compras por objeto separado por metodos

    static makeMetodosCluster = async () => {

        const debito = await prisma.activity.findMany({ where: { metodo: "Debito" } })
        const credito = await prisma.activity.findMany({ where: { metodo: "Credito" } })
        const entrada = await prisma.activity.findMany({ where: { metodo: "Entrada" } })

        const aux = { "debito": debito, "credito": credito, "entrada": entrada }

        return aux
    }

    //buscar compras por objeto separado por categorias

    static makeCategoriaCluster = async () => {
        const categorias = await CategoriaService.getCategorias()

        const corpo: any = {}

        for (let i of categorias) {
            const aux = await prisma.activity.findMany({
                where: {
                    categoriaId: i.id
                },
                include: {
                    categoria: true
                },
                orderBy: [
                    { createdAt: 'desc' }
                ]
            })
            corpo[i.nome] = aux
        }

        return corpo
    }

    private static _retornoVerifica = (message: string): retorno => {
        const aux = { status: false, message }
        return aux
    }

    private static _verifyActivity = (body: Activity): retorno => {

        const { id, userId, tipo, descricao, valor, metodo, parcelaAtual, parcelaTotal, dataPagamento, categoriaName } = body;

        if (metodo !== "Credito" && metodo !== "Debito" && metodo !== "Entrada" && metodo !== "Dinheiro") {
            return this._retornoVerifica("Metodo incorreto")
        }

        if (metodo === "Credito" && parcelaAtual <= 0 && parcelaTotal <= 0) {
            return this._retornoVerifica("Activity de credito sem o número de parcelas")
        }

        if (metodo === "Debito" && parcelaAtual != 0 && parcelaTotal != 0) {
            return this._retornoVerifica("Activity de Debito com o número de parcelas")
        }

        if (descricao == "") {
            return this._retornoVerifica("Descrição não pode ser vazia")
        }

        if (valor <= 0) {
            return this._retornoVerifica("Valor invalido")
        }

        if (parcelaAtual < 0 && parcelaTotal < parcelaAtual && parcelaTotal < 0) {
            return this._retornoVerifica("Valores do parcelamento incorretos")
        }

        if (tipo === "Income" && categoriaName !== "Income") {
            return this._retornoVerifica("Categoria incorreta para income")
        }

        if (tipo === "Income" && metodo !== "Entrada") {
            return this._retornoVerifica("Categoria incorreta para income")
        }

        return { status: true, body: {} }
    }

    private static _dataCredito = (data: Date, metodo: string): Date => {
        const dp: Date = new Date(data);

        const diaI = dp.getDate();
        const mesI = dp.getMonth() + 1;
        const anoI = dp.getFullYear();

        let diaF: number = diaI;
        let mesF: number = mesI;
        let anoF: number = anoI;

        if (metodo === "Credito") {
            diaF = 15;
            if (diaI >= 10) {
                if (mesI === 12) {
                    mesF = 1;
                    anoF = anoI + 1;
                }
                else {
                    mesF = mesI + 1;
                    anoF = anoI;
                }
            }
        }

        const dataPagamentoFinal = new Date(`${anoF}-${mesF < 10 ? "0" + mesF : mesF}-${diaF < 10 ? "0" + diaF : diaF}T03:00:00.000Z`);

        return dataPagamentoFinal;
    }

    //criar compra
    static createActivity = async (body: { user: string, descricao: string, valor: number, tipo: string, metodo: string, parcelaAtual: number, parcelaTotal: number, dataPagamento: string, categoriaName: string }) => {

        const verify: { status: boolean, message?: string, body?: object } = this._verifyActivity(body as unknown as Activity)

        const { user, descricao, valor, tipo, metodo, parcelaAtual, parcelaTotal, dataPagamento, categoriaName } = body;

        if (!verify.status) {
            return verify
        }

        const dataPagamentoFinal = this._dataCredito(new Date(dataPagamento), metodo);

        const dupli = await prisma.activity.count({
            where: {
                descricao: descricao,
                userId: user,
                valor: valor,
                metodo: metodo,
                categoriaName: categoriaName,
                parcelaAtual: parcelaAtual,
                parcelaTotal: parcelaTotal,
                dataPagamento: dataPagamentoFinal
            }
        })

        if (dupli) {
            verify.status = false;
            verify.message = "Produto identico já existe"
            return verify
        }


        verify["body"] = await prisma.activity.create({
            data: {
                descricao: descricao,
                valor: valor,
                metodo: metodo,
                tipo: tipo,
                parcelaAtual: parcelaAtual,
                parcelaTotal: parcelaTotal,
                dataPagamento: dataPagamentoFinal,
                categoria: {
                    connectOrCreate: {
                        where: { nome: categoriaName },
                        create: { nome: categoriaName }
                    }
                },
                user: {
                    connect: {
                        id: user
                    }
                },
            },
            include: {
                categoria: true,
                user: true
            }
        })
        return verify
    }

    //editar um compra por id

    static updateActivity = async (body:Activity) => {

        const { id, userId, descricao, valor, tipo, metodo, parcelaAtual, parcelaTotal, dataPagamento, categoriaName } = body;

        const data = new Date(dataPagamento)

        const verify: { status: boolean, message?: string, body?: object } = this._verifyActivity(body)

        if (!verify.status) {
            return verify
        }

        verify["body"] = await prisma.activity.update({
            where: {
                id: id
            },
            data: {
                descricao: descricao,
                valor: valor,
                tipo: tipo,
                metodo: metodo,
                parcelaAtual: parcelaAtual,
                parcelaTotal: parcelaTotal,
                dataPagamento: data,
                categoria: {
                    connectOrCreate: {
                        where: { nome: categoriaName },
                        create: { nome: categoriaName }
                    }
                },
                user: {
                    connect: {
                        id: userId
                    }
                },

            }
        })

        return verify

    }

    //deletar uma compra

    static deleteActivity = async (id: string) => {

        const ret = await prisma.activity.delete({
            where: {
                id
            },
            select: {
                descricao: true,
                valor: true,
                dataPagamento: true
            }
        })

        return ret
    }

    // adiantar parcelar futuras
    static adiantaActivities =async (id:string) => {
        const atual:Activity = await prisma.activity.findUnique({
            where:{
                id
            }
        })

        const proximos:Array<Activity> = await prisma.activity.findMany({
            where:{
                dataPagamento:{
                    gt:atual.dataPagamento,
                },
                descricao: atual.descricao,
                categoriaId: atual.categoriaId
            }
        })

        for(let i of proximos){
            i.dataPagamento = atual.dataPagamento;
            this.updateActivity(i);
        }

        
    }

}








