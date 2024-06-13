import CategoriaController from "../../controllers/categoria.controller";
import JWTManager from "../../services/jwt.services";



export default class CategoriaSchema{

    static getAllCategorias = {
        schema: {
            tags: ['Categorias'],
            security: [{ apiKey: [] }],
            operationId: 'Get All',
            summary: "Get All",
            response: {
                200: {
                    type: 'array',
                    items: {
                        type: 'object',
                        properties: {
                            id: { type: 'string' },
                            nome: { type: 'string' }
                        }
                    }
                }
            }
        },
        preHandler: [JWTManager.validateToken],
        handler: CategoriaController.getAllCategorias
    }

    static getFilteredCategorias = {
        schema: {
            tags: ['Categorias'],
            security: [{ apiKey: [] }],
            operationId: 'Get filtered',
            summary: "Get filtered",
            body: {
                type: 'object',
                properties: {
                    nome: { type: 'string' }
                }
            },
            response: {
                200: {
                    type: 'array',
                    items: {
                        type: 'object',
                        properties: {
                            id: { type: 'string' },
                            nome: { type: 'string' }
                        }
                    }
                }
            }
        },
        preHandler: [JWTManager.validateToken],
        handler: CategoriaController.getFilteredCategorias
    }


}