import { FastifyReply, FastifyRequest } from "fastify";
import ActivityServices from "../services/activity.services";
import {z} from 'zod'
import { Activity } from "@prisma/client"

export default class ActivityController{

    private static idSchema = z.object({
        id: z.string()
    })  

    static listAllActivities = async(request:FastifyRequest, reply:FastifyReply) => {
        const retorno = await ActivityServices.listAllActivities()
        reply.status(200).send(retorno)
    }

    static listAllByMonth =async (request:FastifyRequest, reply:FastifyReply) => {

        const body = z.object({
            data: z.string()
        })

        const retorno = await ActivityServices.listAllByMonth(body.parse(request.body).data);

        reply.status(200).send(retorno);
    }

    static getIncomeFromMonth =async (request:FastifyRequest, reply:FastifyReply) => {
        const body = z.object({
            data: z.string()
        })
        
        const f = body.parse(request.body);
        const retorno = await ActivityServices.getFromMonth(body.parse(request.body).data, "Income");
        
        reply.status(200).send({valor: retorno});
    }

    static getExpenseFromMonth =async (request:FastifyRequest, reply:FastifyReply) => {
        const body = z.object({
            data: z.string()
        })
        
        const retorno = await ActivityServices.getFromMonth(body.parse(request.body).data, "Expense");

        reply.status(200).send({valor: retorno});
    }

    static listAllByFilter = async(request:FastifyRequest, reply:FastifyReply)=> {

        const activitiesSchema = z.object({
            descricao: z.string(),
            metodo: z.string(),
            tipo: z.string(),
            categoria: z.string(),
            user: z.string(),
            valor: z.number(),
            dataInicio: z.string(),
            dataFim: z.string()
        })

        const retorno = await ActivityServices.listAllByFilter(activitiesSchema.parse(request.body))
        if(retorno){
            reply.status(200).send(retorno)
        }else{
            reply.status(400).send("Não foi possivel encontrar as despesas")
        }
    }

    static makeMetodosCluster = async (request:FastifyRequest, reply:FastifyReply) => {
        const retorno = await ActivityServices.makeMetodosCluster()
        if(retorno){
            reply.status(200).send(retorno)
        }else{
            reply.status(400).send("Não foi possivel encontrar as despesas")
        }
    }
    static makeCategoriaCluster =async (request:FastifyRequest, reply:FastifyReply) => {
        const retorno = await ActivityServices.makeCategoriaCluster()
        if(retorno){
            reply.status(200).send(retorno)
        }else{
            reply.status(400).send("Não foi possivel encontrar as despesas")
        }
    };

    static createActivity = async(request:FastifyRequest, reply:FastifyReply) => {

        const activitySchema = z.object({
            descricao: z.string(),
            metodo: z.string(),
            tipo: z.string(),
            categoriaName: z.string(),
            user: z.string(),
            valor: z.number(),
            parcelaAtual: z.number(),
            parcelaTotal: z.number(),
            dataPagamento: z.string()
        })

        try{
            const body = activitySchema.parse(request.body)

            if(body.metodo === "Credito"){
                let aux = body;
                for(let i = 1; i <= body.parcelaTotal; i++){
                    aux.parcelaAtual = i;
                    const ret = await ActivityServices.createActivity(aux);
                    
                    let newDate = new Date(aux.dataPagamento);
                    let newD = new Date(newDate.getFullYear(), newDate.getMonth() + 1, newDate.getDate());
                    aux.dataPagamento = newD.toDateString();
                    if(!ret.status){
                        reply.status(400).send(ret.message);
                    }
                }
            }
            else {
                const ret = await ActivityServices.createActivity(body);                
                ret.status? reply.status(201).send(ret): reply.status(200).send(ret)
            }
        }
        catch(e:any){
            reply.status(400).send(e.message)
        }
    }

    static updateActivity = async(request:FastifyRequest, reply:FastifyReply) => {
        const activitySchema = z.object({
            id: z.string(),
            descricao: z.string(),
            metodo: z.string(),
            categoriaName: z.string(),
            userId: z.string(),
            valor: z.number(),
            tipo: z.string(),
            parcelaAtual: z.number(),
            parcelaTotal: z.number(),
            dataPagamento: z.string()
        })

        const body = activitySchema.parse(request.body)

        const retorno = await ActivityServices.updateActivity(body as unknown as Activity)

        if(retorno.status){        
            reply.status(200).send(retorno)
        }
        else{
            reply.status(400).send(retorno)
        }
    }

    static deleteActivity = async(request:FastifyRequest, reply:FastifyReply) => {
        const id = this.idSchema.parse(request.body)
        const retorno = ActivityServices.deleteActivity(id.id) 

        return retorno
    }

    static adiantaActivity =async (request:FastifyRequest, reply:FastifyReply) => {
        ActivityServices.adiantaActivities(this.idSchema.parse(request.body).id);
    }
}


