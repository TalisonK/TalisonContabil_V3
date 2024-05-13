package main

import (
	"log"
	"os"

	"github.com/TalisonK/TalisonContabil/src/config"
	"github.com/TalisonK/TalisonContabil/src/database"
	"github.com/TalisonK/TalisonContabil/src/routes"
	"github.com/gofiber/fiber/v3"
)

func main() {

	err := config.Load()

	if err != nil {
		log.Fatal("Erro ao carregar as configurações")
		os.Exit(2)
	}

	err = database.OpenConnectionLocal()

	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados local")
		os.Exit(2)
	}

	err = database.OpenConnectionCloud()

	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados na nuvem")
		os.Exit(2)
	}

	defer database.CloseConnections()

	app := fiber.New()

	routes.Router(app)

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
