import CategoriaSchema from "./schemas/category.schema";


module.exports = (fastify:any, _:any, done:any) => {

    fastify.decorate('categoria', () => {})

    fastify.get("/", CategoriaSchema.getAllCategorias)
    fastify.post("/filter", CategoriaSchema.getFilteredCategorias)

    done()

}




