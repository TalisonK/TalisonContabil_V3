
import UserSchemas from "./schemas/user.schemas"

module.exports = (fastify:any, _:any, done:any) => {

    fastify.decorate('user', () => {})

    fastify.get("/", UserSchemas.todosUsuarios)
    fastify.post("/", UserSchemas.criaUsuario)
    fastify.put("/", UserSchemas.updateUsuario)
    fastify.post("/login", UserSchemas.loginSchema)

    done()

}