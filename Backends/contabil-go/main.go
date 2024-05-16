package main

import (
	"log"
	"os"

	"github.com/TalisonK/TalisonContabil/src/config"
	"github.com/TalisonK/TalisonContabil/src/database"
	"github.com/TalisonK/TalisonContabil/src/routes"
	"github.com/TalisonK/TalisonContabil/src/util"
	"github.com/gofiber/fiber/v3"
)

func main() {

	err := config.Load()

	if err != nil {
		log.Fatal("Erro ao carregar as configurações, certifique-se de que o arquivo de configuração está correto.")
		util.LogHandler("Erro ao carregar as configurações", err, "main")
		os.Exit(2)
	}

	err = database.OpenConnectionLocal()

	if err != nil {
		util.LogHandler("Erro ao conectar ao banco de dados local", err, "main")
	}

	err = database.OpenConnectionCloud()

	if err != nil {
		util.LogHandler("Erro ao conectar ao banco de dados em nuvem", err, "main")
	}

	defer database.CloseConnections()

	app := fiber.New()

	routes.Router(app)

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
