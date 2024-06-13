import { FastifyReply, FastifyRequest } from "fastify"




export default class GeneralSchemas {

    static ping = {
        schema: {
            tags: ['General'],
            operationId: 'Ping',
            summary: "Ping",
            response: {
                200: {
                    type: 'object',
                    items: {
                        type: 'object',
                        properties: {
                            nome: { type: 'string' }
                        }
                    }
                }
            }
        },
        handler: (request:FastifyRequest, reply: FastifyReply) => {          
            reply.status(200).send("Salve minha maquina")
        }
    }

}