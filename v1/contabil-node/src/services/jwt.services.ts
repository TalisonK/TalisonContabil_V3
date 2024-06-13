import { FastifyReply, FastifyRequest } from "fastify"
const jwt = require('jsonwebtoken')

interface returnToken{
    token:string,
    expires: number
}

const expiresIn = 900;

export default class JWTManager{
    static createToken = (payload:string):returnToken => {
        const token = jwt.sign({userId: payload}, process.env.JWTSECRET, {expiresIn: expiresIn})
        const expires = Math.round((Date.now() / 1000) + expiresIn)
        return {token, expires}
    }
    
    static validateToken = (request:FastifyRequest, reply:FastifyReply, next:any) => {
    
        const jwtToken = String(request.headers["authorization"]);

        const start = jwtToken.substring(0,7);

        if(start !== "Bearer "){
            reply.status(401).send("Token invalido!");
            next();
        }

        const token = jwtToken.substring(7);

        jwt.verify(token, process.env.JWTSECRET, (err:any, decoded:any) => {
            if(err) reply.status(401).send("Token invalido!");
    
            next();
        })
    
    }
}

