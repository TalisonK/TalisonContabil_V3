package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/TalisonK/TalisonContabil/src/auth"
	"github.com/TalisonK/TalisonContabil/src/config"
	"github.com/TalisonK/TalisonContabil/src/database"
	"github.com/TalisonK/TalisonContabil/src/routes"
	"github.com/TalisonK/TalisonContabil/src/util"
)

func main() {

	err := config.Load()

	auth.NewAuth()

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

	// Create a new Chi router

	r := routes.Router()

	http.ListenAndServe(fmt.Sprintf(":%s", config.GetServerPort()), r)

}
