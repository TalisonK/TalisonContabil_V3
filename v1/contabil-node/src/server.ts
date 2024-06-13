const {fastifyBuilder} = require("./routes/index.routes");


const server = async() => {
    const app = await fastifyBuilder()

    app.listen({
        host:'0.0.0.0',
        port: process.env.PORT ? Number(process.env.PORT) : 3333
    }).then( () => {
        console.log("Server running")
    })
}
server()