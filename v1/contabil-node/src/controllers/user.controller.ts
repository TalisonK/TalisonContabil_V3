import { FastifyReply, FastifyRequest } from "fastify";
import {z} from 'zod';
import {allUsers, createUser, updateUser, loginChecker} from "../services/user.services"
import JWTManager from "../services/jwt.services";


export default class UserController {
    static getAllUsers = () => {
        const ret = allUsers();
        return ret;
    }
    
    static createOneUser = async (request:FastifyRequest, reply:FastifyReply) => {
    
        const userCreateSchema = z.object({
            nome: z.string(),
            senha: z.string()
        })
    
        const {nome, senha} = userCreateSchema.parse(request.body)
    
        const retorno = await createUser(nome, senha);
    
        if(retorno === ""){
            reply.status(400).send("O nome de usuário já existe")
        }
        else {
            delete retorno["senha"]
            reply.status(201).send(retorno)
        }
    
        return retorno;
    
    }
    
    static updateOneUser = async (request:FastifyRequest, reply:FastifyReply) => {
    
        const userCreateSchema = z.object({
            id: z.string(),
            nome: z.string(),
            senha: z.string()
        })
    
        const {id, nome, senha} = userCreateSchema.parse(request.body)
    
        const retorno = await updateUser(id, nome, senha)
    
        if(retorno){
            reply.status(201).send(retorno)
        }
        else{
            reply.status(400).send(retorno)
        }
    }
    
    static login = async(request:FastifyRequest, reply:FastifyReply) => {
    
        const loginCreateSchema = z.object({
            nome: z.string(),
            senha: z.string()
        })
    
        const {nome, senha} = loginCreateSchema.parse(request.body)
    
        const retorno = await loginChecker(nome, senha)        
    
        if(retorno === null){
            reply.status(401).send("Informações incorretas")
        }
        else{
            const {token, expires} = JWTManager.createToken(retorno)
            reply.status(200).send({"message":"Autenticado!", "token":token, "expiresIn":expires, "userId": retorno})
        }
    }
    
}



