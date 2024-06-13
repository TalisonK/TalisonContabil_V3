import UserController from "../../controllers/user.controller"
import JWTManager from "../../services/jwt.services"


export default class UserSchemas {

    static todosUsuarios = {
        schema: {
            tags: ['User'],
            security: [{ apiKey: [] }],
            operationId: 'todos Usuarios',
            summary: "Busca todos usuários",
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
        handler: UserController.getAllUsers
    }

    static criaUsuario = {
        schema: {
            tags: ['User'],
            security: [{ apiKey: [] }],
            operationId: 'criar Usuarios',
            summary: "Criar um novo usuário",
            body: {
                type: 'object',
                properties: {
                    nome: { type: 'string' },
                    senha: { type: 'string' },
                }
            },
            response: {
                200: {
                    type: 'object',
                    properties: {
                        id: { type: 'string' },
                        nome: { type: 'string' }
                    }
                }
            }
        },
        preHandler: [JWTManager.validateToken],
        handler: UserController.createOneUser
    }

    static updateUsuario = {
        schema: {
            tags: ['User'],
            security: [{ apiKey: [] }],
            operationId: 'atualizar Usuarios',
            summary: "editar um usuário existente",
            body: {
                type: 'object',
                properties: {
                    id: { type: 'string' },
                    nome: { type: 'string' },
                    senha: { type: 'string' },
                }
            },
            response: {
                200: {
                    type: 'object',
                    properties: {
                        id: { type: 'string' },
                        nome: { type: 'string' }
                    }
                }
            }
        },
        preHandler: [JWTManager.validateToken],
        handler: UserController.updateOneUser
    }

    static loginSchema = {
        schema: {
            tags: ["Login"],
            operationId: "Login",
            body: {
                type: 'object',
                properties: {
                    nome: { type: 'string' },
                    senha: { type: 'string' },
                }
            },
            response: {
                200: {
                    type: 'object',
                    properties: {
                        message:{type: 'string'},
                        token: { type: 'string' },
                        expiresIn: {type:'number'},
                        userId: { type: 'string' }
                    }
                }
            }
        },
        handler: UserController.login
    }
}


