import ActivityController from "../../controllers/activity.controller"
import CategoriaService from "../../services/categoria.services"
import JWTManager from "../../services/jwt.services"

export default class ActivitySchemas {
    private static _body = {
        type: 'array',
        items: {
            type: 'object',
            properties: {
                id: { type: 'string' },
                descricao: { type: 'string' },
                valor: { type: 'number' },
                tipo: { type: 'string' },
                metodo: { type: 'string' },
                categoriaId: { type: 'string' },
                categoriaName: { type: 'string' },
                userId: { type: 'string' },
                parcelaAtual: { type: 'number' },
                parcelaTotal: { type: 'number' },
                dataPagamento: { type: 'string' },
                createdAt: { type: 'string' }

            }
        }
    }

    //buscar todas as compras

    static listAllSchema = {
        schema: {
            tags: ['Activities'],
            security: [{ apiKey: [] }],
            operationId: 'todos activities',
            summary: "Busca todos activities",
            response: {
                200: this._body
            }
        },
        preHandler: [JWTManager.validateToken],
        handler: ActivityController.listAllActivities
    }

    static listAllByMonth = {
        schema: {
            tags: ['Activities'],
            security: [{ apiKey: [] }],
            operationId: 'listar activities por mês',
            summary: "listar activities por mês",
            body: {
                data: { type: 'string' }
            },
            response: {
                200: this._body
            }
        },
        preHandler: [JWTManager.validateToken],
        handler: ActivityController.listAllByMonth
    }

    static getIncomeFromMonth = {
        schema: {
            tags: ['Activities'],
            security: [{ apiKey: [] }],
            operationId: 'get income por mês',
            summary: "get income por mês",
            body: {
                data: { type: 'string' }
            },
            response: {
                200: {
                    type: 'object',
                    properties:{
                        valor: {type: 'string'}
                    }
                }
            }
        },
        preHandler: [JWTManager.validateToken],
        handler: ActivityController.getIncomeFromMonth
    }

    static getExpenseFromMonth = {
        schema: {
            tags: ['Activities'],
            security: [{ apiKey: [] }],
            operationId: 'get Expense por mês',
            summary: "get Expense por mês",
            body: {
                data: { type: 'string' }
            },
            response: {
                200: {
                    type: 'object',
                    properties:{
                        valor: {type: 'string'}
                    }
                }
            }
        },
        preHandler: [JWTManager.validateToken],
        handler: ActivityController.getExpenseFromMonth
    }

    //buscar compras por descrição

    static listAllByFiltroSchema = {
        schema: {
            tags: ['Activities'],
            security: [{ apiKey: [] }],
            operationId: 'listar activities com filtro',
            summary: "listar activities com filtro",
            body: {
                descricao: { type: 'string' },
                metodo: { type: 'string' },
                categoria: { type: 'string' },
                user: { type: 'string' },
                tipo: { type: 'string' },
                valor: { type: 'number' },
                dataInicio: { type: 'string' },
                dataFim: { type: 'string' }
            },
            response: {
                200: this._body
            }
        },
        preHandler: [JWTManager.validateToken],
        handler: ActivityController.listAllByFilter
    }


    //buscar compras por objeto separado por metodos

    static makeMetodosClusterSchema = {
        schema: {
            tags: ['Activities'],
            security: [{ apiKey: [] }],
            operationId: 'metodos Cluster',
            summary: "metodos Cluster",
            response: {
                200: {
                    type: 'object',
                    properties: {
                        debito: this._body,
                        credito: this._body,
                        entrada: this._body
                    }
                }
            }
        },
        preHandler: [JWTManager.validateToken],
        handler: ActivityController.makeMetodosCluster
    }

    //buscar compras por objeto separado por categorias

    static makeCategoriasClusterSchema = {
        schema: {
            tags: ['Activities'],
            security: [{ apiKey: [] }],
            operationId: 'Categoria Cluster',
            summary: "Categoria Cluster",
            response: {
                200: {
                    type: 'object',
                    properties: {
                        Jogo: this._body,
                        Saude: this._body,
                        Alimento: this._body,
                        Computador: this._body,
                        Conta: this._body,
                        Ecommerce: this._body,
                        Estudo: this._body,
                        Helen: this._body,
                        Imposto: this._body,
                        Leitura: this._body,
                        Movel: this._body,
                        Periferico: this._body,
                        Servico: this._body,
                        Streaming: this._body,
                        Telefone: this._body,
                        Utensilio: this._body,
                        Veiculo: this._body,
                        Vestimenta: this._body
                    }
                }
            }
        },
        preHandler: [JWTManager.validateToken],
        handler: ActivityController.makeCategoriaCluster
    }

    //criar compra

    static createActivitySchema = {
        schema: {
            tags: ['Activities'],
            security: [{ apiKey: [] }],
            operationId: 'cria uma novo activity',
            summary: 'cria uma novo activity',
            body: {
                descricao: { type: 'string' },
                metodo: { type: 'string' },
                categoriaName: { type: 'string' },
                user: { type: 'string' },
                valor: { type: 'number' },
                tipo: { type: 'string' },
                parcelaAtual: { type: 'number' },
                parcelaTotal: { type: 'number' },
                dataPagamento: { type: 'string' },
            },
            response: {
                201: {
                    type: 'object',
                    properties: {
                        status: { type: 'boolean' },
                        body: {
                            type: 'object',
                            properties: {
                                id: { type: 'string' },
                                descricao: { type: 'string' },
                                valor: { type: 'number' },
                                tipo: { type: 'string' },
                                metodo: { type: 'string' },
                                categoriaId: { type: 'string' },
                                categoriaName: { type: 'string' },
                                userId: { type: 'string' },
                                parcelaAtual: { type: 'number' },
                                parcelaTotal: { type: 'number' },
                                dataPagamento: { type: 'string' },
                                createdAt: { type: 'string' }

                            }
                        }
                    }
                }
            }
        },
        preHandler: [JWTManager.validateToken],
        handler: ActivityController.createActivity

    }

    //editar um compra por id

    static updateActivityById = {
        schema: {
            tags: ['Activities'],
            security: [{ apiKey: [] }],
            operationId: 'atualiza um activity por id',
            summary: 'atualiza um activity por id',
            body: {
                id: { type: 'string' },
                descricao: { type: 'string' },
                metodo: { type: 'string' },
                categoria: { type: 'string' },
                user: { type: 'string' },
                valor: { type: 'number' },
                tipo: { type: 'string' },
                parcelaAtual: { type: 'number' },
                parcelaTotal: { type: 'number' },
                dataPagamento: { type: 'string' },
            },
            response: {
                200: {
                    type: 'object',
                    properties: {
                        status: { type: 'boolean' },
                        body: {
                            type: 'object',
                            properties: {
                                id: { type: 'string' },
                                descricao: { type: 'string' },
                                valor: { type: 'number' },
                                tipo: { type: 'string' },
                                metodo: { type: 'string' },
                                categoriaId: { type: 'string' },
                                categoriaName: { type: 'string' },
                                userId: { type: 'string' },
                                parcelaAtual: { type: 'number' },
                                parcelaTotal: { type: 'number' },
                                dataPagamento: { type: 'string' },
                                createdAt: { type: 'string' }

                            }
                        }
                    }
                }
            }
        },
        preHandler: [JWTManager.validateToken],
        handler: ActivityController.updateActivity
    }

    //deletar uma compra

    static deleteActivity = {
        schema: {
            tags: ['Activities'],
            security: [{ apiKey: [] }],
            operationId: "deleta um activity",
            summary: "deleta um activity",
            body:{
                type: 'object',
                properties:{
                    id: {type: 'string'}
                }
            },
            response: {
                200: {
                    type:'object',
                    properties:{
                        descricao: { type: 'string' },
                        valor: { type: 'number' },
                        dataPagamento: { type: 'string' }
                    }
                }
            }
        },
        preHandler: [JWTManager.validateToken],
        handler: ActivityController.deleteActivity
    }

    static adiantaActivity = {
        schema: {
            tags: ['Activities'],
            security: [{ apiKey: [] }],
            operationId: "adianta activity",
            summary: "adianta activity",
            body:{
                type: 'object',
                properties:{
                    id: {type: 'string'}
                }
            },
            response: {
                200: {
                    type:'object',
                    properties:{
                        descricao: { type: 'string' },
                        valor: { type: 'number' },
                        dataPagamento: { type: 'string' }
                    }
                }
            }
        },
        preHandler: [JWTManager.validateToken],
        handler: ActivityController.adiantaActivity
    }

}