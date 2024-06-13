import { FastifyReply, FastifyRequest } from "fastify"
import swagger from "@fastify/swagger"
import swaggerUI from "@fastify/swagger-ui"
import { SwaggerTheme } from "swagger-themes"

import GeneralSchemas from "./schemas/general.schema"


export const fastifyBuilder = async() => {

  const Fastify = require("fastify")

  const fastify = Fastify({exposeHeadRoutes: false })
  
  await fastify.register(swagger, {
    swagger: {
      info: {
        title: 'Talison\'s Contabil',
        description: 'Documentação da API',
        version: '0.0.1'
      },
      securityDefinitions: {
        apiKey: {
          type: 'apiKey',
          name: 'authorization',
          in: 'header'
        }
      }
    },
    hideUntagged: true,
    exposeRoute: true,
  })
  
  fastify.register(swaggerUI, {
    theme:{
      css: [
        {filename: 'theme.css', content: new SwaggerTheme('v3').getBuffer('dark')}
      ]
    }
  })
  
  // pinger
  fastify.get("/", GeneralSchemas.ping)
  
  //PLUGGINS
  fastify.register(require('./user.routes'), {prefix: "/user"})
  fastify.register(require('./activity.routes'), {prefix: "/activity"})
  fastify.register(require('./categoria.routes'), {prefix: "/categoria"})

  return fastify
}


