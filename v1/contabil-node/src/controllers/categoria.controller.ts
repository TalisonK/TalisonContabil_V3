import CategoriaService from "../services/categoria.services";
import { FastifyReply, FastifyRequest } from "fastify";
import { z } from 'zod';
export default class CategoriaController {

    

    static getAllCategorias = async (request: FastifyRequest, reply: FastifyReply) => {

        const retorno = await CategoriaService.getCategorias()


        reply.status(200).send(retorno)

    }

    static getFilteredCategorias = async (request: FastifyRequest, reply: FastifyReply) => {

        const categorySchema = z.object({
            nome: z.string()
        })        

        const filtro = categorySchema.parse(request.body);

        const retorno = await CategoriaService.getFilteredCategorias(filtro.nome);

        reply.status(200).send(retorno);
    }

}