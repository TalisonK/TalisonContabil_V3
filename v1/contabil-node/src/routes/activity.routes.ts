import ActivitySchemas from "./schemas/activity.schema";


module.exports = (fastify:any, _:any, done:any) => {

    fastify.decorate('Compra', () => {})

    fastify.get("/", ActivitySchemas.listAllSchema)
    fastify.post("/filtro", ActivitySchemas.listAllByFiltroSchema)
    fastify.get("/metodos", ActivitySchemas.makeMetodosClusterSchema)
    fastify.get("/categorias", ActivitySchemas.makeCategoriasClusterSchema)
    fastify.post("/add", ActivitySchemas.createActivitySchema)
    fastify.put("/",ActivitySchemas.updateActivityById)
    fastify.delete("/",ActivitySchemas.deleteActivity)
    fastify.post("/by-month", ActivitySchemas.listAllByMonth)
    fastify.post("/get-income", ActivitySchemas.getIncomeFromMonth)
    fastify.post("/get-expense", ActivitySchemas.getExpenseFromMonth)
    fastify.post("/adianta", ActivitySchemas.adiantaActivity)

    done()

}

























