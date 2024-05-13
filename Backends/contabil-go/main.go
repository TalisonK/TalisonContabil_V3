package main

import (
	"log"
	"os"

	"github.com/TalisonK/media-storager/src/config"
	"github.com/TalisonK/media-storager/src/database"
	"github.com/TalisonK/media-storager/src/routes"
	"github.com/gofiber/fiber/v3"
)

func main() {

	err := config.Load()

	if err != nil {
		log.Fatal("Erro ao carregar as configurações")
		os.Exit(2)
	}

	err = database.OpenConnection()

	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados")
		os.Exit(2)
	}

	defer database.CloseConnection()

	app := fiber.New()

	routes.Router(app)

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
